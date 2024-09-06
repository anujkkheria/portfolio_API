package common

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// check if prod
	prod := os.Getenv("MONGO_URI")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}