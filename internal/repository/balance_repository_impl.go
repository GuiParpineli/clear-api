package repository

import (
	"clear-api/internal/data"
	"clear-api/internal/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type balanceRepository struct{}

func (r *balanceRepository) GetAllBalances() ([]model.BalanceSheet, error) {
	var balances []model.BalanceSheet
	if err := data.DB.Find(&balances).Error; err != nil {
		return nil, err
	}
	return balances, nil
}

func NewBalanceRepository() BalanceRepository {
	return &balanceRepository{}
}

func (r *balanceRepository) checkBalance(balance *model.BalanceSheet) error {
	if balance.Year != time.Now().Year() {
		balance.Year = time.Now().Year()
	}
	if balance.Month < 1 || balance.Month > 12 {
		return errors.New("month is invalid")
	}
	balance.Status = 0
	return nil
}

func (r *balanceRepository) SaveBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
	err := r.checkBalance(balance)
	if err != nil {
		return nil, err
	}
	return data.DB.Create(&balance), nil
}

func (r *balanceRepository) UpdateBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
	var balanceSaved model.BalanceSheet
	if err := data.DB.Where("id = ?", &balance.ID).First(&balanceSaved).Error; err != nil {
		return nil, err
	}

	newBalance := model.BalanceSheet{
		Month:        balance.Month,
		Year:         balanceSaved.Year,
		Status:       1,
		Company:      balanceSaved.Company,
		Account:      balanceSaved.Account,
		Compositions: append(balanceSaved.Compositions, balance.Compositions...),
	}

	return data.DB.Save(&newBalance), nil
}

func (r *balanceRepository) GetBalanceById(balance *model.BalanceSheet) (*gorm.DB, error) {
	return data.DB.
		Joins("Company").
		Joins("Account").
		Joins("Compositions").
		Where("id = ?", &balance.ID).First(&balance), nil
}

func (r *balanceRepository) GetBalanceByCompanyAndAccount(balance *model.BalanceSheet) (*gorm.DB, error) {
	return data.DB.
		Joins("Company").
		Joins("Account").
		Joins("Compositions").
		Where("company = ? AND account = ?", balance.Company, balance.Account).First(&balance), nil
}

func (r *balanceRepository) CloseBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
	balance.Status = 3
	if err := data.DB.Where("id = ?", &balance.ID).First(&balance).Error; err != nil {
		return nil, err
	}
	return data.DB.Save(&balance), nil
}
