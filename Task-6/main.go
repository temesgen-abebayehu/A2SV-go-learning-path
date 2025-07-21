package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"task_manager/data"
	"task_manager/router"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables
	loadEnv()

	// 2. Initialize MongoDB connection with timeout and retry
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Validate required environment variables
	if mongoURI == "" || dbName == "" || jwtSecret == "" {
		log.Fatal("Missing required environment variables (MONGODB_URI, DB_NAME, JWT_SECRET)")
	}

	// Configure MongoDB connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize DB with retry logic
	var err error
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = data.InitializeDB(ctx, mongoURI, dbName)
		if err == nil {
			break
		}
		log.Printf("Connection attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB after %d attempts: %v", maxRetries, err)
	}
	defer data.CloseDB()

	log.Println("Successfully connected to MongoDB!")

	// 3. Set up router
	r := router.SetupRouter(jwtSecret)

	// 4. Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() {
	// Try loading from current directory first
	err := godotenv.Load()
	if err != nil {
		// Try loading from absolute path
		envPath := filepath.Join(".", ".env")
		absPath, _ := filepath.Abs(envPath)
		err = godotenv.Load(absPath)
		if err != nil {
			log.Printf("Note: No .env file found, using system environment variables")
		}
	}
}