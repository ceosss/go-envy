package envy

import (
	"os"
	"strings"
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
	tests := []struct {
		name        string
		env         string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid environment",
			env:     "test",
			wantErr: false,
		},
		{
			name:        "empty environment name",
			env:         "",
			wantErr:     true,
			errContains: "environment name is empty",
		},
		{
			name:        "invalid characters in environment name",
			env:         "test@123",
			wantErr:     true,
			errContains: "invalid characters",
		},
		{
			name:        "non-existent environment file",
			env:         "nonexistent",
			wantErr:     true,
			errContains: "failed to load environment file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "valid environment" {
				setupTestEnvFiles()
				defer teardownTestEnvFiles()
			}

			err := Load(tt.env)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error message %q does not contain %q", err.Error(), tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				// Verify environment variables are set correctly
				if got := os.Getenv("DB_HOST"); got != "testhost" {
					t.Errorf("Expected DB_HOST to be 'testhost', got '%s'", got)
				}
				if got := os.Getenv("API_KEY"); got != "12345" {
					t.Errorf("Expected API_KEY to be '12345', got '%s'", got)
				}
			}
		})
	}
}
