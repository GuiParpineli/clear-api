package handler

import (
	"clear-api/internal/repository"
	"clear-api/internal/service"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	balanceRepo := repository.NewBalanceRepository()
	balanceService := service.NewBalanceService(balanceRepo)

	r.POST("/balance", balanceService.CreateBalance)
	return r
}
