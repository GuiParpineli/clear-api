package model

import "gorm.io/gorm"

type BalanceSheet struct {
	gorm.Model
	Month        int           `gorm:"type:int;NOT NULL" json:"month"`
	Year         int           `gorm:"type:int;NOT NULL" json:"year"`
	Status       BalanceStatus `gorm:"type:varchar(12);NOT NULL" json:"status"`
	Company      *Company      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
	Account      *Account      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"account"`
	Compositions []Composition `gorm:"many2many:balance_sheet_compositions;" json:"compositions"`
}
