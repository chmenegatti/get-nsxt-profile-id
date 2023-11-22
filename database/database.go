package database

import (
	"log"
	"os"

	"github.com/chmenegatti/get-nsxt-profile-id/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "root:password@tcp(172.20.0.2:3306)/nemesis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect do the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully!")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.Organizations{}, &models.Networks{})

	Database = DbInstance{Db: db}

}
