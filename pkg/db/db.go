package db

import (
	"log"
	"pln/jatim/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func Init(url string) *gorm.DB {
	Database, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Database.AutoMigrate(&models.Users{})
	Database.AutoMigrate(&models.IpAddress{})
	Database.AutoMigrate(&models.Group{})

	return Database
}
