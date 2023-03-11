package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	Http                 *Http
	Logging              *Logging
	Kafka                *Kafka
	WriteKeyToKafkaTopic map[string]string
}

type Producer struct {
	ConfigMap *kafka.ConfigMap
}

type Consumer struct {
	ConfigMap *kafka.ConfigMap
}

type Kafka struct {
	Consumer *Consumer
	Producer *Producer
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
			Kafka: &Kafka{
				Producer: &Producer{
					ConfigMap: &kafka.ConfigMap{
						"bootstrap.servers": "127.0.0.1:9092",
						"group.id":          "test",
						"auto.offset.reset": "earliest",
						"client.id":         "tracker-producer",
						"batch.size":        16384,
						"retries":           5,
						"compression.type":  "gzip",
					}},
			},
			WriteKeyToKafkaTopic: map[string]string{},
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
