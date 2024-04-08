package database

import (
	"fmt"
	models "goshop/internal/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func Connect() {
	var dbConfig models.DatabaseConfig

	dbConfig.Username = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.IP = os.Getenv("DB_IP")
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.NameDB = os.Getenv("DB_NAME")

	destination := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.IP, dbConfig.Port, dbConfig.
		NameDB)
	db, err := gorm.Open(mysql.Open(destination), &gorm.Config{})

	if err != nil {
		panic("Cannot Connect Database")
	}

	DbClient = db
	// Comment Migration, if dont use
	migrationDB(db)
}

func migrationDB(db *gorm.DB) {
	// db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Product{})
}
