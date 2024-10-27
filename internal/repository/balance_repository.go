package repository

import (
	"clear-api/internal/data"
	"clear-api/internal/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

func checkBalance(balance *model.BalanceSheet) error {
	if balance.Year != time.Now().Year() {
		balance.Year = time.Now().Year()
	}
	if balance.Month < 1 || balance.Month > 12 {
		return errors.New("month is invalid")
	}
	balance.Status = 0
	return nil
}

func SaveBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
	err := checkBalance(balance)
	if err != nil {
		return nil, err
	}
	return data.DB.Create(&balance), nil
}

func UpdateBalance(balance *model.BalanceSheet) (*gorm.DB, error) {
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
