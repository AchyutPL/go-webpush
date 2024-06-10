package notification

import (
	"go-webpush/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationService NotificationService
}

func NewNotificationController(notificationService NotificationService) NotificationController {
	return NotificationController{
		NotificationService: notificationService,
	}
}

func (nc *NotificationController) SendNotification(ctx *gin.Context) {
	var notification types.Notification

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := nc.NotificationService.SendNotification(&notification)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}
