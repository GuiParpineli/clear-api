package repository

import (
	"clear-api/internal/model"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	SaveBalance(balance *model.BalanceSheet) (*gorm.DB, error)
	UpdateBalance(balance *model.BalanceSheet) (*gorm.DB, error)
	GetBalanceById(balance *model.BalanceSheet) (*gorm.DB, error)
	GetBalanceByCompanyAndAccount(balance *model.BalanceSheet) (*gorm.DB, error)
	CloseBalance(balance *model.BalanceSheet) (*gorm.DB, error)
}
