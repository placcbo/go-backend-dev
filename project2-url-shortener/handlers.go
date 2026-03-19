package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project2-url-shortener/database"
	"strings"
	"time"

	"gorm.io/gorm"
)

func generateUniqueCode() (string, error) {
	for i := 0; i < 10; i++ {
		code := GenerateCode(6)
		_, err := database.GetURL(code)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code, nil
		}
	}

	return "", fmt.Errorf("could not generate unique code after 10 attempts")
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, "url must start with http:// or https://", http.StatusBadRequest)
		return
	}

	code, err := generateUniqueCode()
	if err != nil {
		http.Error(w, "could not generate short code", http.StatusInternalServerError)
		return
	}

	saved, err := database.SaveURL(code, req.URL)
	if err != nil {
		http.Error(w, "could not save url", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{
		ShortCode:   saved.ShortCode,
		OriginalURL: saved.OriginalURL,
		ShortURL:    "http://localhost:8080/" + saved.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.Error(w, "short code is required", http.StatusBadRequest)
		return
	}

	url, err := database.GetURL(code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "short code not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if err := database.IncrementClicks(code); err != nil {
		fmt.Println("warn: failed to increment clicks:", err)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/stats/")
	if code == "" {
		http.Error(w, "short code is required", http.StatusBadRequest)
		return
	}

	url, err := database.GetStats(code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "short code not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	resp := StatsResponse{
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		Clicks:      url.Clicks,
		CreatedAt:   url.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
