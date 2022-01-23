package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
)

func proxyHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodConnect {
		connectHandler(writer, request)
		return
	}
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	newReq := request.Clone(context.Background())
	newReq.RequestURI = ""
	newHeader := newReq.Header
	newHeader.Del("Proxy-Connection")
	response, err := client.Do(newReq)
	if err != nil {
		log.Println(err)
		return
	}
	writerHeader := writer.Header()
	for k, vv := range response.Header {
		writerHeader[k] = vv
	}
	writer.WriteHeader(response.StatusCode)
	if response.ContentLength > 0 {
		_, err := io.Copy(writer, response.Body)
		if err != nil {
			return
		}
	}
}

func connectHandler(writer http.ResponseWriter, request *http.Request) {
	rAddr, _ := net.ResolveTCPAddr("tcp4", request.RequestURI)
	rConn, err := net.DialTCP("tcp4", nil, rAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(rConn)

	hj, ok := writer.(http.Hijacker)
	if !ok {
		return
	}
	lConn, _, err := hj.Hijack()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err := lConn.Close()
		if err != nil {
			return
		}
	}()

	_, err = lConn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	if err != nil {
		return
	}

	done := make(chan struct{})
	go func() {
		_, err := io.Copy(rConn, lConn)
		if err != nil {
			return
		}
		close(done)
	}()

	_, err = io.Copy(lConn, rConn)
	if err != nil {
		return
	}

	<-done
}

func main() {
	//log.SetFlags(log.Lshortfile)
	err := http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(proxyHandler))
	if err != nil {
		log.Println(err)
		return
	}
}
