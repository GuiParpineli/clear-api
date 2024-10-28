package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name      string   `gorm:"type:varchar(55);NOT NULL" json:"name"`
	Cnpj      string   `gorm:"type:varchar(14);NOT NULL" json:"cnpj"`
	Email     string   `gorm:"type:varchar(55);NOT NULL" json:"email"`
	Phone     string   `gorm:"type:varchar(15);NOT NULL" json:"phone"`
	AddressID uint     `json:"address_id"`
	Address   *Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
}
