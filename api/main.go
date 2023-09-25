package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"redis_url/api/routes"
	"crypto/md5"
	"sort"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis container address
		Password: "",           // No password for local Redis
		DB:       0,            // Default DB
	})

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/shorten", routes.ShortenURL).Methods("POST")
	// r.HandleFunc("/{shortURL}", RedirectURL).Methods("GET")
	// r.HandleFunc("/topdomains", GetTopDomains).Methods("GET")

	// Start the HTTP server
	fmt.Println("Redis URL Shortener Service is running on :8080...")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}



























// // RedirectURL redirects to the original URL for a given short URL
// func RedirectURL(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	shortURL := vars["shortURL"]

// 	longURL, err := redisClient.Get(ctx, shortURL).Result()
// 	if err == redis.Nil {
// 		http.Error(w, "Short URL not found", http.StatusNotFound)
// 		return
// 	} else if err != nil {
// 		http.Error(w, "Failed to fetch URL", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, longURL, http.StatusFound)
// }

// // GetTopDomains returns the top domains that have been shortened the most
// func GetTopDomains(w http.ResponseWriter, r *http.Request) {
// 	keys, err := redisClient.Keys(ctx, "*").Result()
// 	if err != nil {
// 		http.Error(w, "Failed to fetch keys", http.StatusInternalServerError)
// 		return
// 	}

// 	domainCounts := make(map[string]int)
// 	for _, key := range keys {
// 		if strings.HasPrefix(key, "http") {
// 			parts := strings.Split(key, "/")
// 			if len(parts) >= 3 {
// 				domain := parts[2]
// 				domainCounts[domain]++
// 			}
// 		}
// 	}

// 	topDomains := GetTopN(domainCounts, 3)
// 	json.NewEncoder(w).Encode(topDomains)
// }

// // GetTopN returns the top N elements from a map based on their values
// func GetTopN(data map[string]int, n int) map[string]int {
// 	topN := make(map[string]int)
// 	for key, value := range data {
// 		topN[key] = value
// 	}

// 	// Sort the map by values in descending order
// 	sortedKeys := make([]string, 0, len(topN))
// 	for key := range topN {
// 		sortedKeys = append(sortedKeys, key)
// 	}

// 	sort.Slice(sortedKeys, func(i, j int) bool {
// 		return topN[sortedKeys[i]] > topN[sortedKeys[j]]
// 	})

// 	// Keep only the top N elements
// 	if len(sortedKeys) > n {
// 		sortedKeys = sortedKeys[:n]
// 	}

// 	result := make(map[string]int)
// 	for _, key := range sortedKeys {
// 		result[key] = topN[key]
// 	}

// 	return result
// }
