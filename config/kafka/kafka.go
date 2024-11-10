package kafka

import (
	"fmt"
	// "os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type BrokerDetail struct {
	Address string
	Port    string
}

func SetupKafkaProducer() *kafka.Producer {
	// brokerAddress := os.Getenv("KAFKA_BROKER")
	// brokerPort:=os.Getenv("KAFKA_BROKER_PORT")
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "1",
	})
	if err != nil {
		fmt.Println("Kafka producer 配置錯誤")
	}
	return p
}

func SetupKafkaConsumer() *kafka.Consumer {
	// brokerAddress := os.Getenv("KAFKA_BROKER")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("Kafka Consumer 配置錯誤")
	}

	return c
}
