package processimage

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	m1 "myapp/Producer/mongo1"
	t1 "myapp/Producer/templete"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nfnt/resize"
	"go.mongodb.org/mongo-driver/bson"
)

func ProcessProductimage(a int) (bool, error) {
	client, productsCollection, err := m1.MongoConnect("vinay1", "product")
	defer client.Disconnect(context.Background())
	productID := a
	fmt.Println(a)
	var product t1.Product
	err = productsCollection.FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)

	if err != nil {
		fmt.Println(err)
	}
	var images []string
	for _, v := range product.ProductImages {
		images = append(images, v)
	}
	fmt.Println(images)
	var filepath []string
	for _, imageURL := range images {
		outputFilePath, err := DownloadCompressAndStoreImage1(imageURL)
		if err != nil {
			fmt.Printf("Error processing image: %v\n", err)
			continue
		}
		filepath = append(filepath, outputFilePath)

	}
	update := bson.M{"$set": bson.M{"compressed_product_images": filepath}}
	_, err = productsCollection.UpdateOne(context.Background(), bson.M{"_id": productID}, update)
	if err != nil {
		return false, err
	} else {
		return true, nil
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
	outputFilePath := fmt.Sprintf("images/%s_%d.jpg", imageFileName, time.Now().UnixNano())
	// outputFilePath := fmt.Sprintf("images/%s", imageFileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return "", err
	}
	// Update the product with the compressed image paths

	defer outputFile.Close()
	return outputFilePath, nil
}

func DownloadCompressAndStoreImage1(imageURL string) (string, error) {
	// Create a directory to store the images if it doesn't exist
	err := os.MkdirAll("images", os.ModePerm)
	if err != nil {
		return "", err
	}

	// Extract the image file name from the URL
	imageFileName := GetImageFileName(imageURL)
	fmt.Println(imageFileName)
	// Create a file to store the compressed image
	outputFilePath := fmt.Sprintf("images/%s_%d.jpg", imageFileName[:3], time.Now().UnixNano())
	fmt.Println(outputFilePath)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return "", err
	}

	// Make an HTTP GET request to download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}

	// Resize the image to a smaller size (e.g., 800x600 pixels) while maintaining aspect ratio
	resizedImg := resize.Resize(700, 0, img, resize.Lanczos2)

	// Encode the resized image as JPEG and write it to the file
	if err := jpeg.Encode(outputFile, resizedImg, nil); err != nil {
		return "", err
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		// fmt.Println("Error getting current working directory:", err)
		// return
	}
	absolutePath := cwd + string(os.PathSeparator) + outputFilePath

	absolutePath = strings.Replace(absolutePath, "/", `\`, -1)
	fmt.Println(absolutePath)
	defer outputFile.Close()
	return absolutePath, nil
}

func GetImageFileName(imageURL string) string {
	// Split the URL by '/' and get the last part as the file name
	parts := strings.Split(imageURL, "/")
	fileNameWithParams := parts[len(parts)-1]

	// Remove any characters that are not allowed in a filename
	validFileName := make([]rune, 0, len(fileNameWithParams))
	for _, r := range fileNameWithParams {
		switch r {
		case ' ', '?', '&', '=', '+':
			// Skip characters that are not allowed
		default:
			validFileName = append(validFileName, r)
		}
	}
	return string(validFileName)
}
