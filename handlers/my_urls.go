package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"url_shortener/db"
	"url_shortener/middlewares"
	"url_shortener/models"
)

// MyURLResponse represents dashboard response

type MyURLResponse struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}

func MyURLsHandler(w http.ResponseWriter, r *http.Request) {
	// get user email from JWT
	userEmail := middlewares.GetUserEmailFromContext(r)
	if userEmail == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	collection := db.Client.Database("url_shortener").Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, map[string]string{
		"user_email": userEmail,
	})
	if err != nil {
		http.Error(w, "Failed to fetch URLs", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var response []MyURLResponse

	for cursor.Next(ctx) {
		var urlDoc models.URL
		if err := cursor.Decode(&urlDoc); err != nil {
			continue
		}

		response = append(response, MyURLResponse{
			OriginalURL: urlDoc.OriginalURL,
			ShortURL:    "http://localhost:8080/" + urlDoc.ShortCode,
			Clicks:      urlDoc.Clicks,
			CreatedAt:   urlDoc.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
