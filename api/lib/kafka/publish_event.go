package kafka

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

// PublishEvent : publish kafka event
func PublishEvent(topic, key string, message interface{}) (err error) {

	if producer == nil {
		producer = getProducer()
	}

	eventChan := make(chan kafka.Event)
	defer close(eventChan)

	byteMessage, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(byteMessage),
			Key:   []byte(key),
		},
		eventChan,
	)

	e := <-eventChan
	m := e.(*kafka.Message)

	return m.TopicPartition.Error
}

func getProducer() *kafka.Producer {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    os.Getenv("KAFKA_HOST"),
		"default.topic.config": kafka.ConfigMap{"acks": "all"},
	})

	return p
}
