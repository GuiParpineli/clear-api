package model

import "gorm.io/gorm"

type Composition struct {
	gorm.Model
	TaxNoteNumber string       `gorm:"type:int;NOT NULL" json:"taxNoteNumber"`
	EmissionDate  string       `gorm:"type:date;NOT NULL" json:"emissionDate"`
	DueDate       string       `gorm:"type:date;NOT NULL" json:"dueDate"`
	ISS           float32      `gorm:"type:float;NOT NULL" json:"iss"`
	INSS          float32      `gorm:"type:float;NOT NULL" json:"inss"`
	IRRF          float32      `gorm:"type:float;NOT NULL" json:"irrf"`
	CSRF          float32      `gorm:"type:float;NOT NULL" json:"csrf"`
	Credit        float32      `gorm:"type:float;NOT NULL" json:"credit"`
	Debit         float32      `gorm:"type:float;NOT NULL" json:"debit"`
	History       string       `gorm:"type:varchar(255);NOT NULL" json:"history"`
	ResponsibleID uint         `json:"responsible_id"`
	Responsible   *Responsible `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"responsible"`
}
