package handlers

import (
	"context"
	"net/http"
	"time"

	"url_shortener/db"
	"url_shortener/models"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// extract short code from path
	shortCode := r.URL.Path[1:] // remove leading "/"

	if shortCode == "" {
		http.NotFound(w, r)
		return
	}

	collection := db.Client.Database("url_shortener").Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var urlDoc models.URL
	err := collection.FindOne(ctx, map[string]string{
		"short_code": shortCode,
	}).Decode(&urlDoc)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	// increment click count
	collection.UpdateOne(
		ctx,
		map[string]string{"short_code": shortCode},
		map[string]interface{}{
			"$inc": map[string]int{"clicks": 1},
		},
	)

	// redirect to original URL
	http.Redirect(w, r, urlDoc.OriginalURL, http.StatusFound)
}
