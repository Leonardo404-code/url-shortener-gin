package database

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	StoreService = &StorageService{}
	Ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatalf("Error init redis: %v", err)
	}

	StoreService.redisClient = redisClient

	return StoreService
}
