package subscription

import (
	"go-webpush/internal/database"
	"go-webpush/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscriptionController struct {
	SubscriptionService SubscriptionService
}

func NewSubscriptionController(ss SubscriptionService) SubscriptionController {
	return SubscriptionController{
		SubscriptionService: ss,
	}
}

func (sc *SubscriptionController) CreateSubscription(ctx *gin.Context) {
	var sub models.Subscription

	if err := ctx.ShouldBindJSON(&sub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := sc.SubscriptionService.SubscribeUser(&sub)

	log.Println("🚀🚀🚀🚀🚀err", err)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}
func (sc *SubscriptionController) GetAllSubscriptions(ctx *gin.Context) {
	var subscriptions []models.Subscription
	results := database.DB.Find(&subscriptions)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(subscriptions), "data": subscriptions})
}

// func (sc *SubscriptionController) RemoveSubscription(ctx *gin.Context) {

// 	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

// }
