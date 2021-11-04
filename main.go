package main

import (
	"context"
	"log"

	appKafka "github.com/felachek/kafka-test/kafka"
)

const (
	topic         = "topico"
	brokerAddress = "localhost:9092"
	groupId       = "G1"
)

func main() {
	log.Println("Kafka has been started...")

	ctx := context.Background()

	go appKafka.Produce(ctx, brokerAddress, topic)
	appKafka.Consume(ctx, brokerAddress, topic, groupId)

	// go appKafka.StartKafka(ctx, brokerAddress, topic, groupId)
	// time.Sleep(10 * time.Minute)
}
