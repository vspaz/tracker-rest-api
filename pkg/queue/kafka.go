package queue

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

func NewKafkaProducer(producerConfig *kafka.ConfigMap, logger *logrus.Logger) *kafka.Producer {
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		logger.Fatalf("failed to init kafka producer")
	}
	logger.Info("kafka producer is initialized")
	return producer
}
