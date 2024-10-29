package main

import (
	//"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config := LoadConfig()

	db, err := NewDB(config.PostgresDSN)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Create the subscriptions table
	if err := db.CreateTable(); err != nil {
		log.Fatal("Error creating table:", err)
	}

	kafkaProducer, err := NewKafkaProducer(config.KafkaBroker)
	if err != nil {
		log.Fatal("Error creating Kafka producer:", err)
	}
	defer kafkaProducer.Close()

	router := gin.Default()

	// Post : /subscribe
	router.POST("/subscribe", func(c *gin.Context) {
		var sub Subscription
		if err := c.ShouldBindJSON(&sub); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := db.Subscribe(sub); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "subscribed"})
	})

	// post : /notifications/send
	router.POST("/notifications/send", func(c *gin.Context) {
		var notification Notification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		message, err := json.Marshal(notification)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message: " + err.Error()})
			return
		}

		if err := kafkaProducer.Produce(notification.Topic, message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
	})

	// Post : /unsubscribe
	router.POST("/unsubscribe", func(c *gin.Context) {
		var sub Subscription
		if err := c.ShouldBindJSON(&sub); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := db.Unsubscribe(sub.UserID, sub.Topics); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsubscribe: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "unsubscribed"})
	})

	// Get : /subscriptions/:user_id
	router.GET("/subscriptions/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")

		subscriptions, err := db.FetchSubscriptions(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, SubscriptionResponse{UserID: userID, Subscriptions: subscriptions})
	})

	router.Run(":8080")
}
