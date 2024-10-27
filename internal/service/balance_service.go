package service

import (
	"clear-api/internal/model"
	"clear-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBalance(c *gin.Context) {
	var input model.BalanceSheet

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var balance model.BalanceSheet

	_, err := repository.SaveBalance(&balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balance not save."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": balance})
}
