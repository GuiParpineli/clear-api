package repository

import (
	"clear-api/internal/model"
)

type BalanceRepository interface {
	SaveBalance(balance *model.BalanceSheet) error
	UpdateBalance(balance *model.BalanceSheet) error
	GetBalanceById(balance *model.BalanceSheet) (*model.BalanceSheet, error)
	GetBalanceByCompanyAndAccount(balance *model.BalanceSheet) (*model.BalanceSheet, error)
	CloseBalance(balance *model.BalanceSheet) error
	GetAllBalances() ([]*model.BalanceSheet, error)
	ReopenCloseBalance(id uint) error
	AddComposition(id uint, composition []*model.Composition) error
}
