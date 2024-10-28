package handler

import (
	"clear-api/internal/data"
	"clear-api/internal/repository"
	"clear-api/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	data.ConnectDb()
	balanceRepo := repository.NewBalanceRepository()
	balanceService := service.NewBalanceService(balanceRepo)

	r.GET("/balance", balanceService.GetAllBalances)
	r.POST("/balance", balanceService.CreateBalance)
	return r
}
