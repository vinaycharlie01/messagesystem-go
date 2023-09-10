package main

import (
	"fmt"
	"time"

	p1 "myapp/Consumer/processimage"
	k1 "myapp/Producer/kafka"
)

func main() {
	// res, _ := k1.ReceiveLatestFromKafka()
	// fmt.Println(res)
	// p1.ProcessProductimage(res)

	for {
		// Receive the latest message from Kafka
		res, err := k1.ReceiveLatestFromKafka()
		if err != nil {
			fmt.Printf("Error receiving message from Kafka: %v\n", err)
			// Add error handling or retries as needed
			time.Sleep(time.Second * 5) // Sleep for a while before retrying
			continue
		}

		// Process the received message
		p1.ProcessProductimage(res)

		// Sleep for a while before checking for the next message
		time.Sleep(time.Second * 2)
	}

}
