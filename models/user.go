package models

import "time"


type User struct{

	ID string `json:"id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	IsVerified bool `json:"is_verified" bson:"is_verified"`
	CreatedAt time.Time`json:"created_at" bson:"created_at"`
}