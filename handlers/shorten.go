package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"url_shortener/db"
	"url_shortener/models"
	"url_shortener/utils"
	"url_shortener/middlewares"
)

// ShortenRequest represents request body
type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

// ShortenResponse represents response body
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.OriginalURL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// get user email from JWT (NOT from client)
	userEmail := middlewares.GetUserEmailFromContext(r)
	if userEmail == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shortCode, err := utils.GenerateShortCode()
	if err != nil {
		http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
		return
	}

	urlDoc := models.URL{
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
		UserEmail:   userEmail,
		Clicks:      0,
		CreatedAt:   time.Now(),
	}

	collection := db.Client.Database("url_shortener").Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, urlDoc); err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	shortURL := "http://localhost:8080/" + shortCode

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortURL: shortURL,
	})
}

