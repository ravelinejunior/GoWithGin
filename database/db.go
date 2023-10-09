package database

import (
	"log"

	user_model "github.com/ravelinejunior/go_api_gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=127.18.0.2 user=raveline password=senha123 dbname=Go_DB port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error opening connection", err.Error())
	}
	DB.AutoMigrate(&user_model.UserModel{})
}
