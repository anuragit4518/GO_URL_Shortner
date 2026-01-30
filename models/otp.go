package models

import "time"

type OTP struct {
	ID        string    `bson:"_id,omitempty"`
	Email     string    `bson:"email"`
	Code      string    `bson:"code"`
	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
}
