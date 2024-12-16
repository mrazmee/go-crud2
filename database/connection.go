package database

import (
	"fmt"
	"log"

	"github.com/mrazmee/go-crud2/config"
	"github.com/mrazmee/go-crud2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	cfg, err := config.LoadConfig()
	if err != nil{
		panic("Failed to load config")
	}


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&models.ToDo{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err) // Pastikan error migrasi tertangkap
	}

	fmt.Println("Database connection successful")
}