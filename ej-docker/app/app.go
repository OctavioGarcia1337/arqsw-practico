package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func StartUp() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":8081")
}
