package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexloney/kin/apps/api/internal/cache"
	"github.com/alexloney/kin/apps/api/internal/db"
)

type statusResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {

	// Obtain the port that we will host the API on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Open a connection to the database and run migrations
	database, err := db.OpenDB()
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer database.Close()
	if err := db.RunMigrations(database); err != nil {
		log.Fatalf("migrations: %v", err)
	}

	// Open a connection to Redis
	redisClient, err := cache.OpenRedis()
	if err != nil {
		log.Fatalf("redis: %v", err)
	}
	defer redisClient.Close()

	// Set up the HTTP server and routes
	mux := http.NewServeMux()

	// Root endpoint for basic health check - replace in the future
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "kin api is running",
		})
	})
	registerHealthRoutes(mux)

	// Start the HTTP server
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           loggingMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("kin api listening on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(startedAt))
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}