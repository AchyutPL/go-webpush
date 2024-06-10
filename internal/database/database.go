package database

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	database = os.Getenv("DB_DATABASE")
// 	password = os.Getenv("DB_PASSWORD")
// 	username = os.Getenv("DB_USERNAME")
// 	port     = os.Getenv("DB_PORT")
// 	host     = os.Getenv("DB_HOST")
// )

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("🚀 Connected Successfully to the Database")
}
