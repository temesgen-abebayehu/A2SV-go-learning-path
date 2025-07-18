package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client        *mongo.Client
	db            *mongo.Database
	taskCollection *mongo.Collection
)

func InitializeDB(connectionString, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Atlas connection options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Set database and collection
	db = client.Database(dbName)
	taskCollection = db.Collection("tasks")

	log.Println("Successfully connected to MongoDB Atlas!")
	return nil
}

// CloseDB remains the same
func CloseDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}