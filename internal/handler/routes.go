package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//config.ConnectDb()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
