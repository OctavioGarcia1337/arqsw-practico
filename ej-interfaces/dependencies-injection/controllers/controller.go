package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	key := c.Param("key")
	value := serv.Get(key)
	c.string(http.StatusOK, value)
}
