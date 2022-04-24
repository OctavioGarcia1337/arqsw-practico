package main

import (
	"github.com/OctavioGarcia1337/arq-software/ej-auth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login",
		controllers.Login)
	engine.Run(":8080")
}
