package database

import (
	"fmt"
	"log"
	"todo-api/internal/configs"
	"todo-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = connect()
	if err != nil {
		panic("failed to connect db")
	}

	err = migrate()
	if err != nil {
		panic("failed to migrate db")
	}

	log.Println("connect to db success")
}

func connect() (*gorm.DB, error) {
	dbConf := configs.Conf.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbConf.Host, dbConf.User, dbConf.Password, dbConf.DBName, dbConf.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrate() error {
	return DB.AutoMigrate(models.Task{})
}
