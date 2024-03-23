package database

import (
	"fmt"
	"log"
	"os"

	"github.com/zakkaizzatur/golang-dts-final-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = os.Getenv("PGHOST")
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbPort = os.Getenv("PGPORT")
	dbname = os.Getenv("PGDATABASE")
	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	fmt.Println("Sukses koneksi ke database")
	db.Debug().AutoMigrate(models.Base{}, models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
