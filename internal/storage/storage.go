package storage

import (
	"ot/internal/order/dtos"
	dtos2 "ot/internal/outbox/dtos"
)

type Storage interface {
	CreateOrder(dto dtos.CreateOrderDto) int
	DeleteOutbox(id int)
	CreateOutbox(data dtos2.CreateOutboxDto) int
	UpdateStatusOutbox(id int, complete bool) int
}
