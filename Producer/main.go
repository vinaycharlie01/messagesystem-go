package main

import (
	"context"
	"log"

	a1 "myapp/Producer/api"
	m1 "myapp/Producer/mongo1"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB and initialize collections
	client, _, err := m1.MongoConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize the Gin router
	r := gin.Default()

	// Define an endpoint to create a product
	r.POST("/create-product", a1.CreateProduct)

	// Start the Gin server
	r.Run(":8080")
}
