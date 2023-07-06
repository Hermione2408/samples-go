package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password
		DB:       0,                // Redis database number
	})

	// Ping the Redis server to check the connection
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Define a route to get a value from Redis
	router.GET("/api/mykey", func(c *gin.Context) {
		// Get the value of "mykey" from Redis
		value, err := client.Get(context.Background(), "mykey").Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get value from Redis"})
			return
		}
		c.JSON(200, gin.H{"value": value})
	})

	// Define a route to set a value in Redis
	router.POST("/api/mykey", func(c *gin.Context) {
		var request struct {
			Value string `json:"value" binding:"required"`
		}

		// Bind JSON request to struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Set the value of "mykey" in Redis
		err := client.Set(context.Background(), "mykey", request.Value, 0).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to set value in Redis"})
			return
		}

		c.JSON(200, gin.H{"message": "Value set in Redis"})
	})

	// Run the Gin router
	router.Run(":8083")
}
