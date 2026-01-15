package config

import (
	"errors"
	"os"
)

const (
	httpHost = "HTTP_HOST"
	httpPort = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	Host string
	Port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHost)
	if len(host) == 0 {
		return nil, errors.New(httpHost + " is required")
	}

	port := os.Getenv(httpPort)
	if len(port) == 0 {
		return nil, errors.New(httpPort + " is required")
	}

	return &httpConfig{
		Host: host,
		Port: port,
	}, nil
}

func (hc *httpConfig) Address() string {
	return hc.Host + ":" + hc.Port
}
