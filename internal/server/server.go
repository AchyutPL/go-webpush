package server

import (
	"go-webpush/configs"
	"go-webpush/internal/database"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	server *gin.Engine
)

func NewServer() error {

	server = gin.Default()

	database.ConnectToDB()

	// cors configuration
	server.Use(configs.CorsConfiguration)

	// register all routes from routes.go
	RegisterRoutes(server)

	err := server.Run(":8080")

	log.Fatal(err)

	return err
}
