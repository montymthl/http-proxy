package utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Proxy Proxy `yaml:"proxy"`
}

type Proxy struct {
	Hostname string `yaml:"host_name"`
	Port     int    `yaml:"port"`
}

func GetConfig() Config {
	var config = &Config{Proxy{
		Hostname: "127.0.0.1",
		Port:     8080,
	}}
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
	log.Println(config)
	return *config
}
