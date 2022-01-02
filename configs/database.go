package configs

import (
	"fmt"
	"log"
	"os"

	"ecommerce-v1/auth/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var dsn string
	var err error

	env := os.Getenv("ENV")

	if env == "PROD" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		HOST := os.Getenv("DB_HOST")
		USERNAME := os.Getenv("DB_USER")
		PASSWORD := os.Getenv("DB_PASSWORD")
		PORT := os.Getenv("DB_PORT")
		DATABASE := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", USERNAME, PASSWORD, HOST, PORT, DATABASE)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}

	DB = db
	err = migration(db)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Account{}, &entity.Token{}, &entity.TokenType{})
	if err != nil {
		return err
	}
	return nil
}
