package main

import (
	"fmt"
	"log"
	"net/smtp"
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
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{"transactions"}

	err = c.SubscribeTopics(topics, nil)

	if err != nil {
		fmt.Println("error subiscribing topics", err)
	}

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			processMessage(msg)
		}
	}
}

func processMessage(msg *kafka.Message) {
	receiver := &model.Receiver{}

	err := receiver.ParserJson(msg.Value)

	if err != nil {
		fmt.Println(err)
	}

	if *msg.TopicPartition.Topic == "transactions" {

		auth := smtp.PlainAuth("", os.Getenv("smtpUser"), os.Getenv("smtpPass"), os.Getenv("smtpHost"))
		to := []string{receiver.Email}
		body := "Hi " + receiver.Name + ", this is a consumer send you a email, may not be the most beautiful \n but the guy who did it until last week never had whiten a golang code\r\n"

		message := []byte("To: " + receiver.Email + "\r\n" +
			"Subject: Hello from Kafka Consumer \r\n" +
			"\r\n" + body)

		err := smtp.SendMail(os.Getenv("smtpHost")+":"+os.Getenv("smtpPort"), auth, os.Getenv("smtpSender"), to, message)

		if err != nil {
			fmt.Println("error sending email:", err)
		}

	}

}
