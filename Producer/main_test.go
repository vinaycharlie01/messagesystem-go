package main

import (
	"os"
	"testing"
)


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
