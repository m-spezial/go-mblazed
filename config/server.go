package config

import "code.m-spezial.de/M-Spezial/go-mblazed/config/env"

type ServerConfig struct {
	Port  int    `env:"MBLAZED_SERVER_PORT,required"`
	Host  string `env:"MBLAZED_SERVER_HOST,required"`
	Debug bool   `env:"MBLAZED_SERVER_DEBUG,default=false"`
}

func NewServerConfig(host string, port int) *ServerConfig {
	return &ServerConfig{
		Port: port,
		Host: host,
	}
}

// NewServerConfigFromEnv create a new ServerConfig and loads the values from the environment variables
func NewServerConfigFromEnv() (*ServerConfig, error) {
	config := &ServerConfig{}

	err := config.LoadConfigFromEnv()

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (sc *ServerConfig) LoadConfigFromEnv() error {
	if err := env.ReadValues(sc); err != nil {
		return err
	}
	return nil
}
