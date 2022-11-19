package boot

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file")
	} else {
		log.Println(".env loaded")
	}
}
