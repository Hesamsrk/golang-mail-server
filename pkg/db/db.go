package db

import (
	"fmt"
	"log"

	"github.com/Hesamsrk/golang-mail-server/pkg/config"
	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://pg:%s@localhost:9099/staging", config.LocalConfig.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Class{}, &models.Student{}, &models.User{})

	return db
}
