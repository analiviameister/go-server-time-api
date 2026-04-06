package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

type TimeResponse struct {
	ServerTime string `json:"server_time"`
	Cached     bool   `json:"cached"`
}

func main() {
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("erro ao conectar no redis: %v", err)
	}

	http.HandleFunc("/", timeHandler)

	port := getEnv("PORT", "8080")
	log.Printf("API iniciada na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	const cacheKey = "server_time"

	cachedTime, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Horario do servidor (cache 1m): " + cachedTime))
		return
	}

	now := time.Now().Format(time.RFC3339)

	redisClient.Set(ctx, cacheKey, now, time.Minute)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Horario atual do servidor: " + now))
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("erro ao serializar resposta: %v", err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}