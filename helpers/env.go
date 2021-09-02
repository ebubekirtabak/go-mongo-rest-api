package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvironment() {
	envPath := fmt.Sprint(os.Getenv("BASE_PATH"), ".env.local")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}
}
