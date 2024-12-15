package caching

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

// InitRedis initializes the Redis client
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Use default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

// SetCache sets a key-value pair in Redis with a TTL (optional)
func SetCache(key string, value interface{}, ttlSeconds int) error {
	return RedisClient.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

// GetCache retrieves a value from Redis
func GetCache(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// InvalidateCache removes a key from Redis
func InvalidateCache(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
