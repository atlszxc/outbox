package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"ot/internal/storage/postgres"
	"time"
)

type Producer struct {
	P *kafka.Producer
}

func (p *Producer) Listen(topic string) {
	storage := postgres.GetStorage(postgres.CONNECTION_STRING)
	for {
		data := storage.GetCompleteOutbox()
		for _, msg := range data {
			p.SendMessage(topic, msg)
			storage.DeleteOutbox(msg.Id)
		}
		time.Sleep(5 * time.Second)
		//p.P.Flush(1000)
	}
}

func (p *Producer) SendMessage(topic string, data any) {
	v, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	delChan := make(chan kafka.Event)

	err = p.P.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          v,
		},
		delChan,
	)

	if err != nil {
		fmt.Println(err)
	}

	answer := <-delChan
	msg := answer.(*kafka.Message)
	fmt.Println(msg.Value)
}

func GetProducer() *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	})

	if err != nil {
		panic("Kafka bad")
	}

	return &Producer{P: p}
}
