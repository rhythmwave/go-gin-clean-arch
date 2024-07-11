package config

import (
	"log"
	"os"
)

type Config struct {
	Translation Translation
	Storage     Storage
}

type Translation struct {
	Endpoint string
	Key      string
	Region   string
}

type Storage struct {
	Account         string
	Key             string
	InputToken      string
	OutputToken     string
	InputUrl        string
	OutputUrl       string
	AccountUrl      string
	OutputContainer string
	InputContainer  string
}

func LoadConfig() Config {
	// Load environment variables from .env file if they don't exist in the system
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("No .env file found")
	// }

	config := Config{
		Translation{
			Key:      GetEnv("AZURE_TRANSLATOR_KEY", ""),
			Region:   GetEnv("AZURE_RESOURCE_LOCATION", ""),
			Endpoint: GetEnv("AZURE_ENDPOINT", ""),
		},
		Storage{
			Account:         GetEnv("AZURE_STORAGE_ACCOUNT", ""),
			Key:             GetEnv("AZURE_STORAGE_KEY", ""),
			InputToken:      GetEnv("AZURE_STORAGE_SAS_INPUT_TOKEN", ""),
			OutputToken:     GetEnv("AZURE_STORAGE_SAS_OUTPUT_TOKEN", ""),
			InputUrl:        GetEnv("AZURE_STORAGE_SAS_INPUT_URI", ""),
			OutputUrl:       GetEnv("AZURE_STORAGE_SAS_OUTPUT_URI", ""),
			AccountUrl:      GetEnv("AZURE_STORAGE_ACCOUNT_URL", ""),
			OutputContainer: GetEnv("AZURE_STORAGE_OUTPUT_CONTAINER", ""),
			InputContainer:  GetEnv("AZURE_STORAGE_INPUT_CONTAINER", ""),
		},
	}

	if config.Translation.Key == "" || config.Storage.Key == "" {
		log.Fatalf("Missing required environment variables")
	}

	return config
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
