package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config holds all configuration values
type Config struct {
	MongoURI        string
	FirebaseKeyPath string
	JWTSecret       string
}

var (
	instance *Config
	once     sync.Once
)

// LoadConfig loads environment variables from .env file and returns a Config instance
func LoadConfig() *Config {
	once.Do(func() {
		// Load environment variables from .env file in the config directory
		if err := godotenv.Load("../config/.env"); err != nil {
			log.Printf("Warning: Error loading .env file: %v", err)
		}

		instance = &Config{
			MongoURI:        getEnv("MONGO_URI", "mongodb://localhost:27017"),
			FirebaseKeyPath: getEnv("FIREBASE_KEY_PATH", "D:\\Chatify\\db\\firebase\\serviceAccountKey.json"),
			JWTSecret:       getEnv("JWT_SECRET", "baf9146e2078c18a6c70afadd9c69762f9ca65803c913e4d9fd3c5b1fc805a86"),
		}

		// Validate required environment variables
		if instance.MongoURI == "" {
			log.Fatal("❌ MONGO_URI not set in .env file")
		}
		if instance.FirebaseKeyPath == "" {
			log.Fatal("❌ FIREBASE_KEY_PATH not set in .env file")
		}
		if instance.JWTSecret == "" {
			log.Fatal("JWT_SECRET environment variable not set")
		}
	})

	return instance
}

// GetConfig returns the singleton Config instance
func GetConfig() *Config {
	if instance == nil {
		return LoadConfig()
	}
	return instance
}

// Helper function to get environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetMongoURI returns the MongoDB connection URI
func GetMongoURI() string {
	return GetConfig().MongoURI
}

// GetFirebaseKeyPath returns the Firebase service account key path
func GetFirebaseKeyPath() string {
	return GetConfig().FirebaseKeyPath
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return GetConfig().JWTSecret
}

// GetJWTSecretBytes returns the JWT secret key as byte array
func GetJWTSecretBytes() []byte {
	return []byte(GetConfig().JWTSecret)
}
