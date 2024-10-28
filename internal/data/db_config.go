package data

import (
	"clear-api/internal/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDb() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DBURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName,
	)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to datase", DbHost)
		log.Fatal("connection:", err)
	} else {
		fmt.Println("We are connected to the database ", DbDriver)
	}

	err = DB.AutoMigrate(
		&model.Address{},
		&model.Account{},
		&model.Company{},
		&model.Responsible{},
		&model.Composition{},
		&model.BalanceSheet{},
	)
	if err != nil {
		return
	}
}
