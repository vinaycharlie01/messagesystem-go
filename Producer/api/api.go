package api

import (
	"context"
	"log"
	"math/rand"
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
	_, err := m1.Collection.InsertOne(context.Background(), product)
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
	CreateUser(product.UserID)
	// defer conn.Close()
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

func RandomName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := make([]byte, length)
	for i := range name {
		name[i] = charset[rand.Intn(len(charset))]
	}
	return string(name)
}

func RandomInts(a int) int {
	return rand.Intn(a) + 5 // Generates a random integer between 6 and 15 (exclusive)
}

func RandomMobileNumber(length int) string {
	if length <= 0 {
		return ""
	}
	const digits = "0123456789"
	num := make([]byte, length)
	for i := 0; i < length; i++ {
		num[i] = digits[rand.Intn(len(digits))]
	}

	return string(num)
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func CreateUser(ID int) {

	var user t1.User
	Userdata := t1.User{
		ID:        ID,
		Name:      RandomName(RandomInts(10)),
		Mobile:    RandomMobileNumber(10),
		Latitude:  RandomFloat(10, 2999),
		Longitude: RandomFloat(20, 10000),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user = Userdata
	clinet, UserCollection, err := m1.MongoConnect("vinay1", "User")
	_, err = UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	defer clinet.Disconnect(context.Background())

}
