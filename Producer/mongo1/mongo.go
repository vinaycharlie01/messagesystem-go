package mon

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var ProductsCollection *mongo.Collection
var UsersCollection *mongo.Collection

func MongoConnect() (*mongo.Client, *mongo.Collection, error) {
	// Initialize the MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	// Initialize the collections
	ProductsCollection = client.Database("vinay1").Collection("products")
	// usersCollection = client.Database("your-database-name").Collection("users")
	return client, ProductsCollection, nil
}
