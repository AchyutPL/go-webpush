package notification

import (
	"go-webpush/internal/types"
	"go-webpush/pkg/logger"
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

type NotificationResponse struct {
	Message string `json:"message" example:"success"`
}

// @BasePath /api
// PingExample godoc
// @Summary Send Notification to User
// @Schemes
// @Description Send Notification to User
// @Tags notification
// @Param RequestBody body types.Notification true "The notification object"
// @Accept json
// @Produce json
// @Success 200 {object} NotificationResponse
// @Router /notification/send [post]
func (nc *NotificationController) SendNotification(ctx *gin.Context) {
	var notification types.Notification

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Sugar().Infow("Notification is", "notification", notification)

	err := nc.NotificationService.SendNotification(&notification)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}
