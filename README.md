# http-proxy

Http Proxy is a simple http proxy server based on golang.

Features:

- Simple http proxy implemented both regular proxy (RFC 7230) and tunnel mode (connect method)
- Relay upstream proxy

## Tech detail

See snapshots subdirectory.

## Installation

Compile:

`go build proxy.go`

## Running

The server run on localhost with port 8080 by default.

`proxy -h`

`proxy.yml` is the main configure file with detailed comment. You can specify another file by `-c` flag.

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