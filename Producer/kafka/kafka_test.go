package kafka

import (
	"context"
	"os"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestInsertintoKafka(t *testing.T) {
	// Replace this with your Kafka broker address if it's different from localhost:9092.
	brokerAddress := "localhost:9092"
	topic := "my-topic"
	partition := 0

	// You can use a test Kafka broker or a mocking library for Kafka in unit tests.
	// For simplicity, we'll use a test broker here.

	// Connect to the Kafka broker.
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		t.Fatalf("failed to dial leader: %v", err)
	}
	defer conn.Close()

	// Call the function and check for errors.
	p1 := 42 // Replace with your test data
	_, err = InsertintoKafka(p1)
	if err != nil {
		t.Fatalf("failed to insert into Kafka: %v", err)
	}

	// Additional testing logic if needed, e.g., verify messages are written correctly.
}

func TestReceiveLatestFromKafka(t *testing.T) {
	// // Replace this with your Kafka broker address if it's different from localhost:9092.
	// brokerAddress := "localhost:9092"
	// topic := "my-topic"
	// partition := 0

	// // You can use a test Kafka broker or a mocking library for Kafka in unit tests.
	// // For simplicity, we'll use a test broker here.

	// // Create a Kafka writer to produce a test message.
	// conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	// if err != nil {
	// 	t.Fatalf("failed to dial leader: %v", err)
	// }
	// defer conn.Close()

	// messageValue := []byte("123") // Replace with your test data
	// _, err = conn.WriteMessages(kafka.Message{Value: messageValue})
	// if err != nil {
	// 	t.Fatalf("failed to write test message to Kafka: %v", err)
	// }

	// Call the function and check for errors.
	result, err := ReceiveLatestFromKafka()
	if err != nil {
		t.Fatalf("failed to receive from Kafka: %v", err)
	}

	// Check if the received result matches the expected result.
	expectedResult := 42 // Replace with the expected result from your test message
	if result != expectedResult {
		t.Errorf("expected result %d, got %d", expectedResult, result)
	}
}

func TestMain(m *testing.M) {
	// Set up any test-specific configuration here.
	// For example, you can start a test Kafka broker.

	// Run the tests.
	code := m.Run()

	// Clean up any resources here.
	// For example, you can stop the test Kafka broker.

	// Exit with the test result code.
	os.Exit(code)
}
