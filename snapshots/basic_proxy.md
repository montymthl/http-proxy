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