package subscription

import "github.com/gin-gonic/gin"

func RegisterSubscriptionRoutes(rg *gin.RouterGroup) {
	ss := NewSubscriptionService()
	sc := NewSubscriptionController(ss)
	subscriptionRoute := rg.Group("/subscribe")
	subscriptionRoute.POST("/create", sc.CreateSubscription)
	subscriptionRoute.POST("/getall", sc.GetAllSubscriptions)
	// subscriptionRoute.DELETE("/delete", sc.RemoveSubscription)
}
