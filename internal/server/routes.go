package server

import (
	"go-webpush/internal/api/notification"
	"go-webpush/internal/api/subscription"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register all the api routes here
func RegisterRoutes(server *gin.Engine) {
	log.Println("🚀🚀🚀🚀🚀🚀 REGISTER ROUTES 🚀🚀🚀🚀🚀🚀🚀")
	server.GET("/", IndexHandler)

	server.GET("/health", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	basepath := server.Group("/api")

	notification.RegisterNotificationnRoutes(basepath)
	subscription.RegisterSubscriptionRoutes(basepath)
}

func IndexHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
