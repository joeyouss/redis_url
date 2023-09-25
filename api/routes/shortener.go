// contains logic for shortening
// ShortenURL handles shortening a given URL

package routes

import(
	"crypto/md5"
	"strings"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"redis_url/api/database"

)
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	longURL := requestData["url"]
	shortURL, err := GetShortURL(longURL)

	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"short_url": shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetShortURL generates a short URL for a given long URL
func GetShortURL(longURL string) (string, error) {
	shortURL, err := redisClient.Get(ctx, longURL).Result()

	if err == redis.Nil {
		// Key not found, generate a short URL
		shortURL = fmt.Sprintf("/%s", GenerateShortCode(longURL))
		err := redisClient.Set(ctx, longURL, shortURL, 0).Err()

		if err != nil {
			return "", err
		}
	}

	return shortURL, err
}

// GenerateShortCode generates a short code for a given long URL
func GenerateShortCode(longURL string) string {
	// Implement your own logic for generating short codes (e.g., hash-based)
	// For simplicity, we use a hash-based approach here
	hash := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(longURL))))
	return hash[:8] // Use the first 8 characters as the short code
}
