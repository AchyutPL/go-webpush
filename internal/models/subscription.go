package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type AuthKeys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

// Implementing the Valuer interface
func (a AuthKeys) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implementing the Scanner interface
func (a *AuthKeys) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan value for AuthKeys")
	}
	return json.Unmarshal(bytes, a)
}

type Subscription struct {
	gorm.Model
	Endpoint       string   `json:"endpoint" binding:"required"`
	ExpirationTime *int64   `json:"expirationTime"`
	Keys           AuthKeys `json:"keys" binding:"required"`
	Test           string   `json:"test" binding:"required"`
}
