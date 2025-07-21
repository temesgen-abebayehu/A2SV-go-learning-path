package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	db             *mongo.Database
	userCollection *mongo.Collection
)

func InitializeDB(ctx context.Context, connectionString, dbName string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPI).
		SetConnectTimeout(10 * time.Second).
		SetSocketTimeout(30 * time.Second).
		SetMaxPoolSize(50)

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		return err
	}

	db = client.Database(dbName)
	userCollection = db.Collection("users")

	go createUserIndexes(context.Background())

	log.Println("Connected to MongoDB!")
	return nil
}

func createUserIndexes(ctx context.Context) {
	// Example: create unique index on email for users collection
	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := userCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Failed to create user indexes: %v", err)
	}
}

func CloseDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}