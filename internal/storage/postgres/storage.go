package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	dtos2 "ot/internal/order/dtos"
	"ot/internal/outbox/dtos"
	"ot/internal/outbox/entity"
	"ot/pkg/logger"
)

const (
	CONNECTION_STRING = "postgres://root:postgres@postgres:5432/order"
)

type PGStorage struct {
	Conn   *pgxpool.Pool
	logger *logger.Logger
}

func (pgs *PGStorage) CreateOrder(dto dtos2.CreateOrderDto) int {
	q := `
		INSERT INTO public."order" (products)
		VALUES (
			$1
		)
		RETURNING id
	`
	var id int
	err := pgs.Conn.QueryRow(context.TODO(), q, &dto.Products).Scan(&id)
	if err != nil {
		pgs.logger.Log(err.Error(), 1)
		panic(err)
	}
	return id
}

func (pgs *PGStorage) DeleteOutbox(id int) {
	q := `
		DELETE FROM outbox
		WHERE id = $1
	`

	_, err := pgs.Conn.Query(context.TODO(), q, &id)
	if err != nil {
		pgs.logger.Log(err.Error(), 1)
		panic(err)
	}

}

func (pgs *PGStorage) CreateOutbox(data dtos.CreateOutboxDto) int {
	fmt.Println(data.OrderId)
	defer pgs.Conn.Close()
	q := `
		INSERT INTO public.outbox(order_id, complete)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int

	err := pgs.Conn.QueryRow(context.TODO(), q, &data.OrderId, &data.Complete).Scan(&id)
	if err != nil {
		pgs.logger.Log(err.Error(), logger.Panic)
		panic(err)
	}
	return id
}
func (pgs *PGStorage) UpdateStatusOutbox(id int, complete bool) int {
	defer pgs.Conn.Close()

	fmt.Println(id, complete)

	q := `
		UPDATE outbox
		SET complete = $1
		WHERE id = $2::integer
		RETURNING id
	`
	var otId int
	err := pgs.Conn.QueryRow(context.TODO(), q, &complete, &id).Scan(&otId)
	if err != nil {
		pgs.logger.Log(err.Error(), logger.Panic)
		panic(err)
	}
	return otId
}

func (pgs *PGStorage) GetCompleteOutbox() []entity.Outbox {
	q := `
		SELECT id, complete
		FROM outbox
		WHERE complete = true
	`

	rows, err := pgs.Conn.Query(context.TODO(), q)
	if err != nil {
		pgs.logger.Log(err.Error(), logger.Panic)
		panic(err)
	}

	defer rows.Close()

	var data []entity.Outbox

	for rows.Next() {
		var item entity.Outbox
		err := rows.Scan(&item.Id, &item.Complete)
		if err != nil {
			pgs.logger.Log(err.Error(), logger.Panic)
			panic(err)
		}
		data = append(data, item)
	}

	return data
}

func GetStorage(connStr string) *PGStorage {
	conn, err := pgxpool.New(context.TODO(), connStr)
	if err != nil {
		panic("Can not connect to database")
	}

	return &PGStorage{
		Conn: conn,
		logger: &logger.Logger{
			Tag: "Postgres",
		},
	}
}
