package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, brokerAddress string, topic string, groupId string) {
	l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupId,
		Logger:  l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		fmt.Println("received: ", string(msg.Value))
	}
}
