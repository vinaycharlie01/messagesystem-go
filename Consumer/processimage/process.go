package processimage

import (
	"context"
	"fmt"
	m1 "myapp/Producer/mongo1"
	t1 "myapp/Producer/templete"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func ProcessProductimage(a int) {
	client, productsCollection, err := m1.MongoConnect()
	defer client.Disconnect(context.Background())
	productID := a
	fmt.Println(a)
	var product t1.Product
	err = productsCollection.FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)

	if err != nil {
		fmt.Println(err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data"})
		// return
	}
	var images []string
	for _, v := range product.ProductImages {
		images = append(images, v)
	}
	fmt.Println(images)
	var filepath []string
	for _, imageURL := range images {
		outputFilePath, err := DownloadCompressAndStoreImage(imageURL)
		if err != nil {
			fmt.Printf("Error processing image: %v\n", err)
			continue
		}
		filepath = append(filepath, "/D/Golangp1/Producer/"+outputFilePath)

	}
	update := bson.M{"$set": bson.M{"compressed_product_images": filepath}}
	_, err = productsCollection.UpdateOne(context.Background(), bson.M{"_id": productID}, update)
	if err != nil {
		fmt.Println(err)
	}

}

func DownloadCompressAndStoreImage(imageURL string) (string, error) {
	// Create a directory to store the images if it doesn't exist
	err := os.MkdirAll("images", os.ModePerm)
	if err != nil {
		return "", err
	}
	// Extract the image file name from the URL
	imageFileName := GetImageFileName(imageURL)

	// Create a file to store the compressed image
	outputFilePath := fmt.Sprintf("images/%s", imageFileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return "", err
	}
	// Update the product with the compressed image paths

	defer outputFile.Close()
	return outputFilePath, nil
}

func GetImageFileName(imageURL string) string {
	// Split the URL by '/' and get the last part as the file name
	parts := strings.Split(imageURL, "/")
	return parts[len(parts)-1]
}
