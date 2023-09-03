package api

import (
	"context"
	"log"
	k1 "myapp/Producer/kafka"
	m1 "myapp/Producer/mongo1"
	t1 "myapp/Producer/templete"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	// Parse JSON request body into a Product struct
	var product t1.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Generate ObjectID for the product and user
	product.ID = product.UserID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	// Insert the product data into the MongoDB collection
	_, err := m1.ProductsCollection.InsertOne(context.Background(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product data"})
		return
	}

	// Insert the user data into the MongoDB collection
	// user := t1.User{
	// 	ID:   product.ID,
	// 	Name: product.ProductName, // You can change this to the actual user name.
	// }
	// _, err = usersCollection.InsertOne(context.Background(), user)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user data"})
	// 	return
	// }
	// Return the created product as JSON
	_, err = k1.InsertintoKafka(product.ID)

	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}
