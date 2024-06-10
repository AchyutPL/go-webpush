package main

import (
	"go-webpush/internal/database"
	"go-webpush/internal/models"
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	database.ConnectToDB()
	err := database.DB.AutoMigrate(&models.Subscription{})

	if err != nil {
		log.Fatal("Migration Failed")
	}
	log.Println("✨Migration Successful✨")
}
