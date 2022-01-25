package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Upstream Upstream `yaml:"upstream"`
}

type Server struct {
	Hostname string `yaml:"host_name"`
	Port     int    `yaml:"port"`
}

type Upstream struct {
	Enabled  bool   `yaml:"enabled"`
	Scheme   string `yaml:"scheme"`
	Hostname string `yaml:"host_name"`
	Port     int    `yaml:"port"`
}

func GetConfig(configFile string) Config {
	var config = getDefaultConfig()
	if fp, err := os.Open(configFile); err == nil {
		defer func(fp *os.File) {
			err := fp.Close()
			if err != nil {
				log.Println(err)
			}
		}(fp)
		err := yaml.NewDecoder(fp).Decode(config)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return *config
}

func GetHttpClient(config Config) *http.Client {
	var client = http.DefaultClient
	if config.Upstream.Enabled {
		var uConfig = config.Upstream
		var upstream = fmt.Sprintf("%s://%s:%d", uConfig.Scheme, uConfig.Hostname, uConfig.Port)
		parsed, err := url.Parse(upstream)
		if err != nil {
			log.Println(err)
			return nil
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(parsed)}
	}
	return client
}

func getDefaultConfig() *Config {
	var defaultServer = Server{
		Hostname: "127.0.0.1",
		Port:     8080,
	}
	var defaultUpstream = Upstream{
		Enabled:  false,
		Scheme:   "http",
		Hostname: "127.0.0.1",
		Port:     8081,
	}
	var config = &Config{
		Server:   defaultServer,
		Upstream: defaultUpstream,
	}
	return config
}
