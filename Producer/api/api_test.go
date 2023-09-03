package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	m1 "myapp/Producer/mongo1"
	t1 "myapp/Producer/templete"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestServer() *gin.Engine {
	// Initialize the MongoDB client and collections for testing
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	m1.ProductsCollection = client.Database("vinay1").Collection("products") // Use a test database

	// Initialize the Gin router
	r := gin.Default()

	// Define an endpoint to create a product
	r.POST("/create-product", CreateProduct)

	return r
}

func TestCreateProductEndpoint(t *testing.T) {
	// Setup a test Gin server
	r := setupTestServer()

	// Define the test data for creating a product
	productData := t1.Product{
		// UserID:            1,
		ID:                 1,
		ProductName:        "Test Product",
		ProductDescription: "This is a test product",
		ProductImages:      []string{"test-image1.jpg", "test-image2.jpg"},
		ProductPrice:       19.99,
	}

	// Convert product data to JSON
	productJSON, err := json.Marshal(productData)
	if err != nil {
		t.Fatalf("Failed to marshal product data: %v", err)
	}

	// Create a POST request with the JSON payload
	req, err := http.NewRequest("POST", "/create-product", bytes.NewReader(productJSON))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Parse the response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Check if the response contains the expected message
	if message, ok := response["message"].(string); ok {
		if message != "Product created successfully" {
			t.Errorf("Expected message 'Product created successfully', got '%s'", message)
		}
	} else {
		t.Errorf("Expected 'message' field in the response")
	}

	// Optionally, you can perform additional checks to verify the data in the database.
}
