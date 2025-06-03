package envy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

// ErrInvalidEnvName is returned when the environment name is empty or contains invalid characters
var ErrInvalidEnvName = errors.New("environment name cannot be empty and must only contain alphanumeric characters and underscores")

// Load loads a .env.{env} file like .env.dev or .env.prod
// Returns an error if the environment name is invalid or if the file cannot be loaded
func Load(env string) error {
	// Validate environment name
	if env == "" {
		return fmt.Errorf("%w: environment name is empty", ErrInvalidEnvName)
	}

	// Check for invalid characters in environment name
	if strings.ContainsAny(env, "!@#$%^&*()+=[]{}|\\:;\"'<>,.?/~`") {
		return fmt.Errorf("%w: environment name contains invalid characters", ErrInvalidEnvName)
	}

	filename := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("failed to load environment file %s: %w", filename, err)
	}

	fmt.Printf("go-envy: loaded environment %s\n", env)
	return nil
}
