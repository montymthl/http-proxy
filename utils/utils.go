package utils

import (
	"gopkg.in/yaml.v2"
	"log"
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
	Hostname string `yaml:"host_name"`
	Port     int    `yaml:"port"`
}

func GetConfig() Config {
	var config = getDefaultConfig()
	if fp, err := os.Open("proxy.yml"); err == nil {
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
	}
	return *config
}

func GetHttpClient() {
	
}

func getDefaultConfig() *Config {
	var defaultServer = Server{
		Hostname: "127.0.0.1",
		Port:     8080,
	}
	var defaultUpstream = Upstream{
		Enabled:  false,
		Hostname: "",
		Port:     0,
	}
	var config = &Config{
		Server:   defaultServer,
		Upstream: defaultUpstream,
	}
	return config
}
