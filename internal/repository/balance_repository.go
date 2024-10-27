package repository

import (
	"clear-api/internal/config"
	"clear-api/internal/model"
	"gorm.io/gorm"
)

func SaveBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
	return config.DB.Create(&balance), nil
}
