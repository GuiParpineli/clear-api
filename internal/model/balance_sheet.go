package model

import "gorm.io/gorm"

type BalanceSheet struct {
	gorm.Model
	Month        int           `gorm:"type:int;NOT NULL" json:"month"`
	Year         int           `gorm:"type:int;NOT NULL" json:"year"`
	Status       BalanceStatus `gorm:"type:int;NOT NULL" json:"status"`
	CompanyID    uint          `json:"company_id"`
	Company      *Company      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
	AccountID    uint          `json:"account_id"`
	Account      *Account      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"account"`
	Compositions []Composition `gorm:"many2many:balance_sheet_compositions;" json:"compositions"`
}
