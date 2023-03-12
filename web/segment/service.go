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

func (d *DefaultSegmentService) Enrich(event handlers.Event) {

}

func (d *DefaultSegmentService) ExtractWorkspaceId(event handlers.Event) {

}

func (d *DefaultSegmentService) IsWriteKeyValid(writeKey string) bool {
	return false
}

func (d *DefaultSegmentService) SaveBatch(batch handlers.EventBatch) {

}

func (d *DefaultSegmentService) Save(batch handlers.EventBatch) {

}

func (d *DefaultSegmentService) Preprocess(event handlers.Event, parentEvent handlers.Event) {
	
}