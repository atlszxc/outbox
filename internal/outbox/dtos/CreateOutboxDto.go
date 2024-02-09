package dtos

type CreateOutboxDto struct {
	OrderId  int `json:"order_id"`
	Complete bool
}
