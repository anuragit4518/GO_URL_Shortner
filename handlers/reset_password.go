package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"url_shortener/db"
	"url_shortener/models"
	"url_shortener/utils"
)

// ResetPasswordRequest represents reset password payload

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OTP         string `json:"otp"`
	NewPassword string `json:"new_password"`
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ResetPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.OTP == "" || req.NewPassword == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	otpCollection := db.Client.Database("url_shortener").Collection("otps")

	var otpDoc models.OTP
	err = otpCollection.FindOne(ctx, map[string]string{
		"email": req.Email,
		"code":  req.OTP,
	}).Decode(&otpDoc)

	if err != nil {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	if time.Now().After(otpDoc.ExpiresAt) {
		http.Error(w, "OTP expired", http.StatusUnauthorized)
		return
	}

	// hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	userCollection := db.Client.Database("url_shortener").Collection("users")

	_, err = userCollection.UpdateOne(
		ctx,
		map[string]string{"email": req.Email},
		map[string]interface{}{
			"$set": map[string]string{
				"password": hashedPassword,
			},
		},
	)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	// delete OTP after use
	otpCollection.DeleteOne(ctx, map[string]string{
		"email": req.Email,
		"code":  req.OTP,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password reset successful"))
}
