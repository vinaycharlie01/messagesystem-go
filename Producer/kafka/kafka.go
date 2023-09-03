package kafka

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func InsertintoKafka(p1 int) (*kafka.Conn, error) {
	// to produce messages
	messageValue := []byte(strconv.Itoa(p1))
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	// conn.
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(messageValue)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return conn, nil
}

func ReceiveLatestFromKafka() (int, error) {
	topic := "my-topic"
	partition := 0
	// Create a Kafka reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB minimum batch size
		MaxBytes:  1e6,  // 1MB maximum batch size
	})

	defer r.Close()

	// Seek to the end of the partition to get the latest message
	r.SetOffset(kafka.LastOffset)

	message, err := r.ReadMessage(context.Background())
	if err != nil {
		return 0, err
	}

	// fmt.Println(string(message.Value))
	res, err := strconv.Atoi(string(message.Value))
	if err != nil {
		return 0, err
	}

	return res, nil
}
