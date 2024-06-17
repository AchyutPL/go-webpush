package notification

import (
	"encoding/json"
	"go-webpush/internal/database"
	"go-webpush/internal/models"
	"go-webpush/internal/types"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/SherClockHolmes/webpush-go"
	"gorm.io/gorm"
)

type NotificationService struct {
	DB *gorm.DB
}

func NewNotificationService() NotificationService {
	return NotificationService{
		DB: database.DB,
	}
}

var (
	vapidPublicKey  = "BNsA0W3MD-sMpzk68tisDNa4Dzljur-9hBCsMFco5y_JZ16pyNFGkxiQai1IojcQOFdesb_HFgvfH9cN999WdrU"
	vapidPrivateKey = "DDxVcREL9HU80xycKG4Mtir2hllq8eM8m_QTIeHP-OI"
	mu              sync.Mutex
)

func (ns *NotificationService) SendNotification(notification *types.Notification) error {

	payload := types.Notification{
		Title: notification.Title,
		Body:  notification.Body,
	}

	updatedPayload := map[string]string{
		"title": payload.Title,
		"body":  payload.Body,
	}

	payloadBytes, payloadErr := json.Marshal(updatedPayload)

	if payloadErr != nil {
		return payloadErr
	}

	mu.Lock()
	defer mu.Unlock()

	var wg sync.WaitGroup

	var subscriptions []models.Subscription

	// get all subscriptions
	dbResponse := ns.DB.Find(&subscriptions)

	if dbResponse.Error != nil {
		// panic(dbResponse.Error)
		return dbResponse.Error
	}
	// map over all the subscriptions and send notification concurrently
	for _, sub := range subscriptions {
		sb := &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				Auth:   sub.Keys.Auth,
				P256dh: sub.Keys.P256dh,
			},
		}
		wg.Add(1)
		go func(sub models.Subscription) {
			defer wg.Done()
			resp, err := webpush.SendNotification(payloadBytes, sb, &webpush.Options{
				Subscriber:      "mailto:achyutpaudel24@gmail.com",
				VAPIDPublicKey:  vapidPublicKey,
				VAPIDPrivateKey: vapidPrivateKey,
			})
			if err != nil {
				log.Printf("💥💥💥Error sending notification:💥💥💥 %v\n", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				log.Printf("Non-OK response from push service: %s\n", string(bodyBytes))
			}
		}(sub)
	}
	wg.Wait()

	return dbResponse.Error
}
