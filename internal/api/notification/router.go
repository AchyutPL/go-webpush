package notification

import (
	"github.com/gin-gonic/gin"
)

func RegisterNotificationnRoutes(rg *gin.RouterGroup) {
	ns := NewNotificationService()
	nc := NewNotificationController(ns)

	NotificationRoute := rg.Group("/notification")
	NotificationRoute.POST("/send", nc.SendNotification)
}
