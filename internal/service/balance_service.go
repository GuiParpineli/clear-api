package service

import (
	"clear-api/internal/model"
	"clear-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BalanceService struct {
	repo repository.BalanceRepository
}

func NewBalanceService(repo repository.BalanceRepository) *BalanceService {
	return &BalanceService{repo: repo}
}

func (service *BalanceService) GetAllBalances(c *gin.Context) {
	balances, err := service.repo.GetAllBalances()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balances not found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": balances})
}

func (service *BalanceService) CreateBalance(c *gin.Context) {
	var input model.BalanceSheet

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := service.repo.SaveBalance(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balance not saved."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}
