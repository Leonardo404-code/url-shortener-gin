package database

import (
	"log"
)

func SaveUrl(shortUrl, originalUrl, userId string) {

	err := StoreService.redisClient.Set(Ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		log.Fatalf("Failed to save url: %v", err)
	}
}

func RetrieveOriginalUrl(shortUrl string) string {
	result, err := StoreService.redisClient.Get(Ctx, shortUrl).Result()

	if err != nil {
		log.Fatalf("Failed to retrieve url: %v", err)
	}

	return result
}
