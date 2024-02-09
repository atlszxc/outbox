package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ot/internal/order/dtos"
	"ot/internal/storage/postgres"
)

func CreateOrderHandler(ctx *gin.Context) {
	var data dtos.CreateOrderDto
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Bad request data",
		})
	}

	storage := postgres.GetStorage(postgres.CONNECTION_STRING)
	id := storage.CreateOrder(data)
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
