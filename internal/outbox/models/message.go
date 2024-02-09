package models

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Message struct {
	OutBoxId     int
	KafkaMessage kafka.Message
	Result       bool
	Retry        bool
	Error        error
}
