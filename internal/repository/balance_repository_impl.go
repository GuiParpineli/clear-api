package repository

import (
	"clear-api/internal/data"
	"clear-api/internal/model"
	"errors"
	"time"
)

type balanceRepository struct{}

func NewBalanceRepository() BalanceRepository {
	return &balanceRepository{}
}

func (r *balanceRepository) AddComposition(id uint, composition []*model.Composition) error {
	err := data.DB.Where("id=? AND status != ?", id, 1).Association("Compositions").Append(composition)
	if err != nil {
		return err
	}
	return nil
}

func (r *balanceRepository) ReopenCloseBalance(id uint) error {
	var balance model.BalanceSheet
	err := data.DB.Where("id = ? AND status != ?", id, 2).First(&balance).Error
	if err != nil {
		return err
	}

	if balance.Compositions[0].Responsible.Role != 0 {
		return errors.New("user is not authorized")
	}

	balance.Status = 1
	data.DB.Updates(&balance)
	return nil
}

func (r *balanceRepository) GetAllBalances() ([]*model.BalanceSheet, error) {
	var balances []*model.BalanceSheet
	if err := data.DB.
		Joins("Company").
		Joins("Account").
		Preload("Compositions").
		Find(&balances).Error; err != nil {
		return nil, err
	}
	return balances, nil
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

func (r *balanceRepository) SaveBalance(balance *model.BalanceSheet) error {
	err := r.checkBalance(balance)
	if err != nil {
		return err
	}
	data.DB.Create(&balance)
	return nil
}

func (r *balanceRepository) UpdateBalance(balance *model.BalanceSheet) error {
	var balanceSaved model.BalanceSheet
	if err := data.DB.Where("id = ?", &balance.ID).First(&balanceSaved).Error; err != nil {
		return err
	}

	newBalance := model.BalanceSheet{
		Month:        balance.Month,
		Year:         balanceSaved.Year,
		Status:       1,
		Company:      balanceSaved.Company,
		Account:      balanceSaved.Account,
		Compositions: append(balanceSaved.Compositions, balance.Compositions...),
	}

	data.DB.Save(&newBalance)
	return nil
}

func (r *balanceRepository) GetBalanceById(balance *model.BalanceSheet) (*model.BalanceSheet, error) {
	err := data.DB.
		Joins("Company").
		Joins("Account").
		Preload("Compositions").
		Where("id = ?", &balance.ID).First(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *balanceRepository) GetBalanceByCompanyAndAccount(balance *model.BalanceSheet) (*model.BalanceSheet, error) {
	err := data.DB.
		Joins("Company").
		Joins("Account").
		Preload("Compositions").
		Where("company = ? AND account = ?", balance.Company, balance.Account).First(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *balanceRepository) CloseBalance(balance *model.BalanceSheet) error {
	balance.Status = 3
	if err := data.DB.Where("id = ?", &balance.ID).First(&balance).Error; err != nil {
		return err
	}
	data.DB.Save(&balance)
	return nil
}
