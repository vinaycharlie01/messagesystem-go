package main

import (
	"fmt"

	p1 "myapp/Consumer/processimage"
	k1 "myapp/Producer/kafka"
)

// // MongoDB connection variables
// var client *mongo.Client
// var productsCollection *mongo.Collection
// var usersCollection *mongo.Collection

// func ReceiveLatestFromKafka() (int, error) {
// 	topic := "my-topic"
// 	partition := 0
// 	// Create a Kafka reader
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{"localhost:9092"},
// 		Topic:     topic,
// 		Partition: partition,
// 		MinBytes:  10e3, // 10KB minimum batch size
// 		MaxBytes:  1e6,  // 1MB maximum batch size
// 	})

// 	defer r.Close()

// 	// Seek to the end of the partition to get the latest message
// 	r.SetOffset(kafka.LastOffset)

// 	message, err := r.ReadMessage(context.Background())
// 	if err != nil {
// 		return 0, err
// 	}

// 	// fmt.Println(string(message.Value))
// 	res, err := strconv.Atoi(string(message.Value))
// 	if err != nil {
// 		return 0, err
// 	}

// 	return res, nil
// }

// func copyFile(src, dst string) error {
// 	input, err := ioutil.ReadFile(src)
// 	if err != nil {
// 		return err
// 	}
// 	return ioutil.WriteFile(dst, input, 0644)
// }

// func MongoConnect() (*mongo.Client, *mongo.Collection, error) {
// 	// Initialize the MongoDB client
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	// Initialize the collections
// 	productsCollection = client.Database("vinay1").Collection("products")
// 	// usersCollection = client.Database("your-database-name").Collection("users")

// 	return client, productsCollection, nil
// }

// func processProductimage(a int) {
// 	client, productsCollection, err := MongoConnect()
// 	defer client.Disconnect(context.Background())
// 	productID := a
// 	fmt.Println(a)
// 	var product t1.Product
// 	err = productsCollection.FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)

// 	if err != nil {
// 		fmt.Println(err)
// 		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data"})
// 		// return
// 	}
// 	var images []string
// 	for _, v := range product.ProductImages {
// 		images = append(images, v)
// 	}
// 	fmt.Println(images)
// 	var filepath []string
// 	for _, imageURL := range images {
// 		outputFilePath, err := downloadCompressAndStoreImage(imageURL)
// 		if err != nil {
// 			fmt.Printf("Error processing image: %v\n", err)
// 			continue
// 		}
// 		filepath = append(filepath, "/D/Golangp1/Producer/"+outputFilePath)

// 	}
// 	update := bson.M{"$set": bson.M{"compressed_product_images": filepath}}
// 	_, err = productsCollection.UpdateOne(context.Background(), bson.M{"_id": productID}, update)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }

// func downloadCompressAndStoreImage(imageURL string) (string, error) {
// 	// Create a directory to store the images if it doesn't exist
// 	err := os.MkdirAll("images", os.ModePerm)
// 	if err != nil {
// 		return "", err
// 	}
// 	// Extract the image file name from the URL
// 	imageFileName := getImageFileName(imageURL)

// 	// Create a file to store the compressed image
// 	outputFilePath := fmt.Sprintf("images/%s", imageFileName)
// 	outputFile, err := os.Create(outputFilePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	// Update the product with the compressed image paths

// 	defer outputFile.Close()
// 	return outputFilePath, nil
// }

// func getImageFileName(imageURL string) string {
// 	// Split the URL by '/' and get the last part as the file name
// 	parts := strings.Split(imageURL, "/")
// 	return parts[len(parts)-1]
// }

func main() {
	res, _ := k1.ReceiveLatestFromKafka()
	fmt.Println(res)
	p1.ProcessProductimage(res)

}
