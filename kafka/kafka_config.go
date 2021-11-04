package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartKafka(ctx context.Context, brokerAddress string, topic string, groupId string) {
	conf := kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		GroupID:  groupId,
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Some error occured: ", err)
			continue
		}
		log.Println("Message is : ", string(m.Value))

	}
}
