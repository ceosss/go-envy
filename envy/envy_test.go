package envy

import (
	"os"
	"testing"
)

// Setup: create dummy env files before tests run
func setupTestEnvFiles() {
	os.WriteFile(".env.test", []byte("DB_HOST=testhost\nAPI_KEY=12345"), 0644)
}

// Teardown: remove env files after tests
func teardownTestEnvFiles() {
	os.Remove(".env.test")
}

func TestLoad(t *testing.T) {
	setupTestEnvFiles()
	defer teardownTestEnvFiles()

	Load("test")

	if got := os.Getenv("DB_HOST"); got != "testhost" {
		t.Errorf("Expected DB_HOST to be 'testhost', got '%s'", got)
	}

	if got := os.Getenv("API_KEY"); got != "12345" {
		t.Errorf("Expected API_KEY to be '12345', got '%s'", got)
	}
}
