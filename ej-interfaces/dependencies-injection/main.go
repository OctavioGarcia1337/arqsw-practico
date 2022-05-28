package main

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviogarcia1337/arq-software/ej-interfaces/dependencies-injection/controllers"
)

func main() {
	engine := gin.New()
	engine.GET("/:key", controllers.Get)
	engine.POST("/:key/:value", controllers.Save)
	engine.Run(":8080")
}
