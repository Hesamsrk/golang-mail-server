package db

import (
	"log"

	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:password@localhost:9099/staging"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Class{}, &models.Student{},&models.User{})

	return db
}
