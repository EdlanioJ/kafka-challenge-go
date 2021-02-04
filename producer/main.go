package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/edlanioj/kafka-challenge-go/model"
	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	deliveryChan := make(chan kafka.Event)
	producer, err := newProducer()

	if err != nil {
		fmt.Println("error creaating producer", err)
	}

	receiver := &model.Receiver{
		Email: "pedro@gmail.com",
		Name:  "Pedro Francisco",
	}
	out, err := receiver.ToJson()

	if err != nil {
		fmt.Println(err)
	}

	err = publish(string(out), "transactions", producer, deliveryChan)

	if err != nil {
		fmt.Println("error creaating producer", err)
	}

	deliveryReport(deliveryChan)
}

func newProducer() (*kafka.Producer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func publish(msg string, topic string, producer *kafka.Producer, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

func deliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to", ev.TopicPartition)
			}
		}
	}
}
