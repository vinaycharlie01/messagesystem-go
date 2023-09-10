package mon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoConnect(t *testing.T) {
	databaseName := "vinay1"
	collectionName := "product"

	client, collection, err := MongoConnect(databaseName, collectionName)
	if err != nil {
		t.Fatalf("MongoConnect returned an error: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Verify that the client and collection are not nil
	assert.NotEqual(t, client, nil)
	assert.NotEqual(t, collection, nil)
	if client == nil {
		t.Fatal("MongoConnect returned a nil client")
	}
	if collection == nil {
		t.Fatal("MongoConnect returned a nil collection")
	}

	// Verify that the collection name matches the expected value
	if collection.Name() != collectionName {
		t.Fatalf("Expected collection name: %s, Got: %s", collectionName, collection.Name())
	}

	if collection.Database().Name() != databaseName {
		t.Fatalf("Expected collection name: %s, Got: %s", collectionName, collection.Name())
	}

	// // Verify that the collections are of type *mongo.Collection
	// _, isCollection := C
	// assert.True(t, isCollection)
	// _, isUsersCollection := UsersCollection.(*mongo.Collection)
	// assert.True(t, isUsersCollection)

	// Verify that the client is connected
	err = client.Ping(context.Background(), nil)
	assert.NoError(t, err)

	// Close the client connection when done
	err = client.Disconnect(context.Background())
	assert.NoError(t, err)

	// Insert your additional tests here as needed
}
