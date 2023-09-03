// kafka/producer_test.go

package kafka

import (
	"log"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/mock"
)

var conn *kafka.Conn

type MockKafkaConnection struct {
	mock.Mock
}

func (m *MockKafkaConnection) SetWriteDeadline(time.Time) error {
	args := m.Called(time.Now())
	return args.Error(0)
}

func (m *MockKafkaConnection) WriteMessages(messages ...kafka.Message) (int, error) {
	args := m.Called(messages)
	return args.Int(0), args.Error(1)
}

func (m *MockKafkaConnection) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestInsertintoKafka(t *testing.T) {
	// Create a mock Kafka connection
	mockConn := new(MockKafkaConnection)
	productID := 123
	conn1, err := InsertintoKafka(productID)
	if err != nil {
		log.Fatal(err)
	}
	conn = conn1
	// Define the test product ID
	// productID := 123 // Example product ID

	// Expectations for the mock Kafka connection
	mockConn.On("SetWriteDeadline", mock.AnythingOfType("time.Time")).Return(nil)
	mockConn.On("WriteMessages", mock.AnythingOfType("[]kafka.Message")).Return(1, nil)
	mockConn.On("Close").Return(nil)

	// Call the function to insert into Kafka, passing the mock connection
	InsertintoKafka(productID)
	// Assert that the expectations were met
	mockConn.AssertExpectations(t)

	// Assert that there is no error returned
	// if err != nil {
	// 	t.Errorf("Expected no error, got: %v", err)
	// }
}

func TestReceiveLatestFromKafka(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			want:    29,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReceiveLatestFromKafka()
			if err != nil && got != 28 {
				t.Error("error in kafka")
			}
		})
	}
}
