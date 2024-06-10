package subscription

import (
	"go-webpush/internal/database"
	"go-webpush/internal/models"

	"gorm.io/gorm"
)

type SubscriptionService struct {
	DB *gorm.DB
}

func NewSubscriptionService() SubscriptionService {
	return SubscriptionService{
		DB: database.DB,
	}
}

// Save the subscription object to DB
func (ss *SubscriptionService) SubscribeUser(subscription *models.Subscription) error {

	dbResponse := ss.DB.Create(&subscription)

	return dbResponse.Error
}

// Remove the subscription object to DB
// func (ss *SubscriptionService) UnSubscribeUser(subscription *models.Subscription) error {
// 	dbResponse := ss.DB.Delete(&subscription)

// 	return dbResponse.Error
// }
