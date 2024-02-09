package entity

type Outbox struct {
	Id       int
	OrderId  int
	Complete bool
}
