package mon

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection
var UsersCollection *mongo.Collection

func MongoConnect(database string, collection string) (*mongo.Client, *mongo.Collection, error) {
	// Initialize the MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	// Initialize the collections
	Collection = client.Database(database).Collection(collection)
	UsersCollection = client.Database(database).Collection("users")
	return client, Collection, nil
}
