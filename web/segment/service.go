package segment

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/tracker-rest-api/web/handlers"
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

func (d *DefaultSegmentService) Enrich(message handlers.Message) {

}

func (d *DefaultSegmentService) ExtractWorkspaceId(message handlers.Message) {

}

func (d *DefaultSegmentService) IsWriteKeyValid(writeKey string) bool {
	return false
}

func (d *DefaultSegmentService) SaveBatch(batch handlers.Batch) {

}

func (d *DefaultSegmentService) Save(batch handlers.Batch) {

}

func (d *DefaultSegmentService) Preprocess(message handlers.Message, parentMessage handlers.Message) {

}
