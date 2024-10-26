package model

import "gorm.io/gorm"

type AccountType int

const (
	CLIENT AccountType = iota
	SUPPLIER
	EMPLOYEE
)

type Account struct {
	gorm.Model
	Number uint16      `gorm:"type:int;NOT NULL" json:"number"`
	Name   string      `gorm:"type:varchar(55);NOT NULL" json:"name"`
	Type   AccountType `gorm:"type:int;NOT NULL" json:"type"`
}
