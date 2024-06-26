package mongodb

import (
	"context"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func MongoDBSetup() {
	// MongoDB connection string
	mongoURI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Connected to MongoDB!")

	MongoDB = client
}
