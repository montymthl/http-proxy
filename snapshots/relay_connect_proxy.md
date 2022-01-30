```shell
http-proxy>curl -v -x 127.0.0.1:8080 https://api.myip.la
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
* allocate connect buffer!
* Establish HTTP proxy tunnel to api.myip.la:443
> CONNECT api.myip.la:443 HTTP/1.1
> Host: api.myip.la:443
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
> Host: api.myip.la
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Server: nginx
< Date: Sun, 30 Jan 2022 09:13:57 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 13
< Connection: keep-alive
< Keep-Alive: timeout=20
<
x.x.x.x* Connection #0 to host 127.0.0.1 left intact
```