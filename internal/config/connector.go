package config

import (
	"fmt"
	"golang-test-task-betera/internal/model"
	"golang-test-task-betera/internal/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	config := GetConfigurationInstance()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName, config.DbSslMode,
	)
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(status.DbConnectionError().Error(), config)
	}

	err = db.AutoMigrate(&model.Apod{})
	if err != nil {
		log.Fatalf("error migration " + err.Error())
	}
}

func GetDBInstance() *gorm.DB {
	return db
}
