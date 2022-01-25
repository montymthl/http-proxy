package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/montymthl/http-proxy/utils"
	"io"
	"log"
	"net"
	"net/http"
)

var configFile = "proxy.yml"
var config utils.Config

func proxyHandler(writer http.ResponseWriter, request *http.Request) {
	//log.Println(request)
	if request.Method == http.MethodConnect {
		connectHandler(writer, request)
		return
	}
	client := utils.GetHttpClient(config)
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
	flag.StringVar(&configFile, "c", "proxy.yml", "Specify the main config file")
	flag.Parse()
	config = utils.GetConfig(configFile)
	var addr = fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port)
	log.Println("server started on:" + addr)
	err := http.ListenAndServe(addr, http.HandlerFunc(proxyHandler))
	if err != nil {
		log.Println(err)
		return
	}
}
