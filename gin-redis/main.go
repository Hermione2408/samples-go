package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	router := gin.Default()

	router.GET("/api/mykey", getValue)
	router.POST("/api/mykey", setValue)
	router.DELETE("/api/mykey", deleteValue)
	router.PUT("/api/mykey", updateValue)

	router.Run(":8083")
}

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Print("Failed to connect to Redis:", err)
	}

	return client
}

func getValue(c *gin.Context) {
	client := getRedisClient()
	value, err := client.Get(context.Background(), "mykey").Result()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get value from Redis"})
		return
	}
	c.JSON(200, gin.H{"value": value})
}

func setValue(c *gin.Context) {
	client := getRedisClient()
	var request struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := client.Set(context.Background(), "mykey", request.Value, 0).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to set value in Redis"})
		return
	}

	c.JSON(200, gin.H{"message": "Value set in Redis"})
}

func deleteValue(c *gin.Context) {
	client := getRedisClient()
	err := client.Del(context.Background(), "mykey").Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete value from Redis"})
		return
	}

	c.JSON(200, gin.H{"message": "Value deleted from Redis"})
}

func updateValue(c *gin.Context) {
	client := getRedisClient()
	var request struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := client.Set(context.Background(), "mykey", request.Value, 0).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update value in Redis"})
		return
	}

	c.JSON(200, gin.H{"message": "Value updated in Redis"})
}
