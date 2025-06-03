package envy

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// Load loads a .env.{env} file like .env.dev or .env.prod
func Load(env string) {
	filename := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("go-envy: failed to load %s: %v", filename, err)
	}
	fmt.Printf("go-envy: loaded environment %s\n", env)
}
