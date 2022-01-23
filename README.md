# http-proxy

Http Proxy is a simple http proxy server based on golang.

Features:

- Simple http proxy implemented both regular proxy (RFC 7230) and tunnel mode (connect method)

## Packet sniffer

Regular proxy:

```shell
http-proxy>curl -v -x http://127.0.0.1:8080 http://apple.com
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET http://apple.com/ HTTP/1.1
> Host: apple.com
> User-Agent: curl/7.79.1
> Accept: */*
> Proxy-Connection: Keep-Alive
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 301 Moved Permanently
< Cache-Control: no-store
< Cdnuuid: bb9cc9e3-0147-4b34-bcb3-530f2b8a98ea-2400913451
< Content-Language: en
< Content-Length: 304
< Content-Type: text/html
< Date: Sun, 23 Jan 2022 12:37:16 GMT
< Location: https://www.apple.com/
< Server: ATS/9.0.3
< Via: http/1.1 hkhkg3-edge-bx-006.ts.apple.com (ApacheTrafficServer/9.0.3)
< X-Cache: none
<
<HTML>
<HEAD>
<TITLE>Document Has Moved</TITLE>
</HEAD>

<BODY BGCOLOR="white" FGCOLOR="black">
<H1>Document Has Moved</H1>
<HR>

<FONT FACE="Helvetica,Arial"><B>
Description: The document you requested has moved to a new location.  The new location is "https://www.apple.com/".
</B></FONT>
<HR>
</BODY>
* Connection #0 to host 127.0.0.1 left intact
```
Tunnel mode:

```shell
http-proxy>curl -v -x http://127.0.0.1:8080 -p https://apple.com
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
* allocate connect buffer!
* Establish HTTP proxy tunnel to apple.com:443
> CONNECT apple.com:443 HTTP/1.1
> Host: apple.com:443
> User-Agent: curl/7.79.1
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 Connection established
<
* Proxy replied 200 to CONNECT request
* CONNECT phase completed!
* schannel: disabled automatic use of client certificate
* schannel: ALPN, offering http/1.1
* schannel: ALPN, server accepted to use http/1.1
> GET / HTTP/1.1
> Host: apple.com
> User-Agent: curl/7.79.1
> Accept: */*
>
* schannel: server closed the connection
* Mark bundle as not supporting multiuse
< HTTP/1.1 301 Redirect
< Date: Sun, 23 Jan 2022 12:42:21 GMT
< Connection: close
< Via: http/1.1 hkhkg3-edge-bx-006.ts.apple.com (ApacheTrafficServer/9.0.3)
< Server: ATS/9.0.3
< Cache-Control: no-store
< Location: https://www.apple.com/
< Content-Type: text/html
< Content-Language: en
< CDNUUID: bb9cc9e3-0147-4b34-bcb3-530f2b8a98ea-2401743271
< X-Cache: none
< Content-Length: 304
<
<HTML>
<HEAD>
<TITLE>Document Has Moved</TITLE>
</HEAD>

<BODY BGCOLOR="white" FGCOLOR="black">
<H1>Document Has Moved</H1>
<HR>

<FONT FACE="Helvetica,Arial"><B>
Description: The document you requested has moved to a new location.  The new location is "https://www.apple.com/".
</B></FONT>
<HR>
</BODY>
* Closing connection 0
* schannel: shutting down SSL/TLS connection with apple.com port 443
```
## Installation

Compile:

`go build proxy.go`

## Running

The server run on localhost with port 8080 by default.

`./proxy`

## Using

On Browser, you can install extension SwitchyOmega and config proxy server localhost:8080 with type http.

SwitchyOmega extension use tunnel mode.

On command line, you can set http_proxy/https_proxy environment with `http://server_domain:8080`

Windows：

```shell
set http_proxy=http://server_domain:8080
set https_proxy=http://server_domain:8080
```

Unix/Linux/MacOS:

```shell
export http_proxy=http://server_domain:8080
export https_proxy=http://server_domain:8080
```