package server

import (
	"go-webpush/configs"
	"go-webpush/docs"
	"go-webpush/internal/database"
	"go-webpush/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine
)

func NewServer() error {

	server = gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	// Use the logger
	logger.InitializeLogger()

	database.ConnectToDB()

	// cors configuration
	server.Use(configs.CorsConfiguration)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// register all routes from routes.go
	RegisterRoutes(server)

	err := server.Run(":8080")

	log.Fatal(err)

	return err
}
