package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {
	// configuration
	kafkaBrokers := []string{"localhost:9092"}        // localhost:9092 as per the kafka cluster as defined in the /golang/message-bus/kafka/hello-kafka/docker-compose.yml
	version, err := sarama.ParseKafkaVersion("2.1.1") // version of the kafka server
	if err != nil {
		panic(err)
	}
	kafkaGroup := "test-group-123"      // unique group name
	kafkaTopics := []string{"topic123"} // note: to create the topic see the setup instructions in readme.md

	// set kafka configuration
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = version
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	// 		OffsetOldest to start processing from the start
	// 		OffsetNewest to start processing from the new messages onwards

	// create kafka consumer group
	consumerGroupHandler := handler{
		ready: make(chan bool),
	}
	ctx, cancel := context.WithCancel(context.Background())
	consumerGroup, err := sarama.NewConsumerGroup(kafkaBrokers, kafkaGroup, kafkaConfig)
	if err != nil {
		panic(err)
	}

	// wait for the kafka consumer group to be ready
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, kafkaTopics, &consumerGroupHandler); err != nil {
				panic(err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumerGroupHandler.ready = make(chan bool)
		}
	}()

	<-consumerGroupHandler.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = consumerGroup.Close(); err != nil {
		panic(err)
	}
}

// handler implements the ConsumerGroupHandler interface
type handler struct {
	ready chan bool
}

// Setup method implementation
func (h *handler) Setup(sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

// Cleanup method implementation
func (h *handler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim method implementation
func (h *handler) ConsumeClaim(s sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received: offset = %d: %v", message.Offset, string(message.Value))

		//
		// do: process the message here
		//

		// mark the message as read
		s.MarkMessage(message, "")
	}
	return nil
}
