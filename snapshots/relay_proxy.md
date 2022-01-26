```shell
http-proxy>curl -v -x 127.0.0.1:8080 http://api.myip.la
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET http://api.myip.la/ HTTP/1.1
> Host: api.myip.la
> User-Agent: curl/7.79.1
> Accept: */*
> Proxy-Connection: Keep-Alive
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Connection: keep-alive
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
< Date: Wed, 26 Jan 2022 02:23:11 GMT
< Server: nginx
<
x.x.x.x* Connection #0 to host 127.0.0.1 left intact
```