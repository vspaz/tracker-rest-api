package config

import (
	"sync"
	"time"
)

var (
	config *SingletonConfig
	once   sync.Once
)

type SingletonConfig struct {
	Config *Conf
}

type Conf struct {
	Http    *Http
	Logging *Logging
}

type Http struct {
	*Server
}

type Server struct {
	HostAndPort             string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	IdleTimeout             time.Duration
	RequestExecutionTimeout time.Duration
}

type Logging struct {
	Level string
}

func initConfig() *SingletonConfig {
	return &SingletonConfig{
		Config: &Conf{
			Http: &Http{
				Server: &Server{
					HostAndPort:             ":8080",
					ReadTimeout:             10 * time.Second,
					WriteTimeout:            10 * time.Second,
					IdleTimeout:             10 * time.Second,
					RequestExecutionTimeout: 10 * time.Second,
				},
			},
			Logging: &Logging{
				Level: "info",
			},
		},
	}
}

func GetConfig() *SingletonConfig {
	once.Do(
		func() {
			config = initConfig()
		},
	)
	return config
}
