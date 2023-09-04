package main

import (
	"fmt"
	"math/rand"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func randomName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := make([]byte, length)
	for i := range name {
		name[i] = charset[rand.Intn(len(charset))]
	}
	return string(name)
}

func main() {
	for i := 0; i < 5; i++ {
		name := randomName(8) // Generate a random name with 8 characters
		fmt.Println(name)
	}
}
