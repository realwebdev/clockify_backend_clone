package datastore

import (
	"log"

	"github.com/jinzhu/gorm"
)

func AutoMigrate(modelObject interface{}, db *gorm.DB) error {
	if err := db.DropTableIfExists(modelObject).Error; err != nil {
		log.Print("Error occured While Deleting Previous Table")
		return err
	}
	log.Print("Previous Table drop successfully")

	if err := db.AutoMigrate(modelObject).Error; err != nil {
		return err
	}
	log.Print("models created successfully")

	return nil
}
