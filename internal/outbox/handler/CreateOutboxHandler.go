package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ot/internal/outbox/dtos"
	"ot/internal/storage/postgres"
)

func CreateOutboxHandler(ctx *gin.Context) {
	var data dtos.CreateOutboxDto
	err := ctx.BindJSON(&data)
	if err != nil {
		panic("zxczxczxc")
	}

	st := postgres.GetStorage(postgres.CONNECTION_STRING)
	id := st.CreateOutbox(data)

	//p := kafka.GetProducer()
	//p.SendMessage("test", data)
	//
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
