package models

import "time"

type URL struct {
	ID          string    `bson:"_id,omitempty"`
	OriginalURL string    `bson:"original_url"`
	ShortCode   string    `bson:"short_code"`
	UserEmail   string    `bson:"user_email"`
	Clicks      int       `bson:"clicks"`
	CreatedAt   time.Time `bson:"created_at"`
}
