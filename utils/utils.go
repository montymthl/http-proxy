// Package utils some useful functions in this project
package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Upstream Upstream `yaml:"upstream"`
	Log      Log      `yaml:"log"`
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

type Log struct {
	Enabled bool   `yaml:"enabled"`
	Level   string `yaml:"level"`
	Output  string `yaml:"output"`
}

func GetConfig(configFile string) Config {
	var config = getDefaultConfig()
	if fp, err := os.Open(configFile); err == nil {
		defer func(fp *os.File) {
			err := fp.Close()
			if err != nil {
				log.Print(err)
			}
		}(fp)
		err := yaml.NewDecoder(fp).Decode(config)
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Print(err)
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
			log.Print(err)
			return nil
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(parsed)}
	}
	return client
}

//Set config default value
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
	var defaultLog = Log{
		Enabled: false,
		Level:   "info",
		Output:  "",
	}
	var config = &Config{
		Server:   defaultServer,
		Upstream: defaultUpstream,
		Log:      defaultLog,
	}
	return config
}

func SetupLog(verbose bool, config Config) {
	var Logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	if verbose {
		Logger = Logger.Level(zerolog.DebugLevel).With().Caller().Logger()
	}
	if config.Log.Enabled {
		level, _ := zerolog.ParseLevel(config.Log.Level)
		Logger = Logger.Level(level)
		if len(config.Log.Output) > 0 {
			fp, err := os.OpenFile(config.Log.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Print(err)
				return
			}
			Logger = Logger.Output(fp)
		}
	}
	log.Logger = Logger
}
