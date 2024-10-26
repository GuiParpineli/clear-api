package model

type Address struct {
	Street  string `gorm:"type:varchar(80);NOT NULL" json:"street"`
	Number  string `gorm:"type:varchar(10);NOT NULL" json:"number"`
	ZipCode string `gorm:"type:varchar(10);NOT NULL" json:"zip_code"`
	City    string `gorm:"type:varchar(55);NOT NULL" json:"city"`
	State   string `gorm:"type:varchar(2);NOT NULL" json:"state"`
}
