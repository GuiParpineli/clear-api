package model

import "gorm.io/gorm"

type Role uint16

const (
	ADMIN Role = iota
	USER
	SUPER_ADMIN
)

type Responsible struct {
	gorm.Model
	Name      string   `gorm:"type:varchar(55);NOT NULL" json:"name"`
	Cpf       string   `gorm:"type:varchar(11);" json:"cpf"`
	Email     string   `gorm:"type:varchar(55);NOT NULL" json:"email"`
	Phone     string   `gorm:"type:varchar(15);NOT NULL" json:"phone"`
	CompanyID uint     `json:"company_id"`
	Company   *Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
	Role      Role     `gorm:"type:int;NOT NULL" json:"role"`
}
