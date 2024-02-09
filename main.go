package main

import (
	"log"
	"ot/internal/outbox/handler"
	"ot/pkg/kafka"
	server2 "ot/pkg/server"
)

func main() {
	server := server2.NewServer(server2.DEBUG)

	producer := kafka.GetProducer()
	go producer.Listen("outbox")

	server.AddRoute(server2.POST, "/create", handler.CreateOutboxHandler)
	server.AddRoute(server2.PATCH, "/update", handler.UpdateStatusOutbox)

	err := server.S.Run(":8080")
	if err != nil {
		log.Fatalln("Server broken")
	}
}
