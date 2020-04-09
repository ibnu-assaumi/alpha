package kafka

import (
	"os"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/stretchr/testify/assert"
)

func TestPublishEvent(t *testing.T) {
	t.Parallel()

	t.Run("ERROR_PUBLISH_EVENT_MESSAGE", func(t *testing.T) {
		assert.Error(t, PublishEvent("test", "test", make(chan bool)))
	})

	t.Run("ERROR_PRODUCE_MESSAGE", func(t *testing.T) {
		os.Setenv("KAFKA_HOST", "foo:9092")
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"message.timeout.ms": 1000,
		})
		assert.NoError(t, err)
		producer = p
		assert.Error(t, PublishEvent("test", "test", "test"))
	})
}

func Test_getProducer(t *testing.T) {
	t.Run("POSITIVE_GET_PRODUCER", func(t *testing.T) {
		assert.NotNil(t, getProducer())
	})
}
