package main

import (
	"github.com/Shopify/sarama"

	"fmt"
)

func main() {
	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true

	kafkaHost := "localhost:9092"
	kafkaTopic := "topic123" // note: topic is created from command line using `./kafka-topics.sh --create ....` command. see the "notes" file for more details

	brokers := []string{kafkaHost}

	// Create new consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	// to get the topics:
	ts, err := consumer.Topics()
	if err != nil {
		fmt.Printf("topics err: ", err)
		return
	}
	fmt.Println(ts)

	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("ConsumePartition err: ", err)
		return
	}

	for m := range partitionConsumer.Messages() {
		fmt.Printf("offset: %d: %s", m.Offset, string(m.Value))
	}

}
