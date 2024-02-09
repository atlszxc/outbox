package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ot/internal/outbox/dtos"
	"ot/internal/storage/postgres"
)

func UpdateStatusOutbox(ctx *gin.Context) {
	var data dtos.UpdateStatusOutboxDto
	ctx.BindJSON(&data)
	fmt.Println(data)
	st := postgres.GetStorage(postgres.CONNECTION_STRING)
	id := st.UpdateStatusOutbox(data.Id, data.Complete)
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
