package database

import (
	"fmt"
	"todo-api/internal/configs"
	"todo-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect() (*gorm.DB, error) {
	dbConf := configs.Conf.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbConf.Host, dbConf.User, dbConf.Password, dbConf.DBName, dbConf.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(models.Task{})
}
