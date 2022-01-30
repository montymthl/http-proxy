package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/montymthl/http-proxy/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"net/http"
)

var configFile = "proxy.yml"
var config utils.Config

func proxyHandler(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
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
		log.Print(err)
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
	rConn, err := utils.GetRemoteConnection(request.RequestURI, config)
	if err != nil {
		log.Print(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}(rConn)

	hj, ok := writer.(http.Hijacker)
	if !ok {
		return
	}
	lConn, _, err := hj.Hijack()
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		err := lConn.Close()
		if err != nil {
			return
		}
	}()

	if !config.Upstream.Enabled {
		_, err = lConn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			return
		}
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
	flag.StringVar(&configFile, "c", "proxy.yml", "Specify the main config file")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Log/Show verbose messages")
	flag.Parse()

	config = utils.GetConfig(configFile)
	utils.SetupLog(verbose, config)

	var addr = fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port)
	log.Info().Msg("Http proxy server started on:" + addr)
	err := http.ListenAndServe(addr, http.HandlerFunc(proxyHandler))
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return
	}
}
