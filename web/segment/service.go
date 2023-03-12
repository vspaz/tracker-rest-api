package segment

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

const (
	ERROR_INVALID_WRITE_KEY = "invalid write key"
	ERROR_EMPTY_BATCH       = "Batch is empty"
	ERROR_MESSAGE_TOO_BIG   = "Event is too big"
	ERROR_UNEXPECTED        = "Unexpected error"
)

type DefaultSegmentService struct {
	Logger          *logrus.Logger
	WriteKeyToTopic map[string]string
	KafkaProducer   *kafka.Producer
}
