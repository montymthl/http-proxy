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