package jobs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EnsureIndexes checks and creates necessary indexes in MongoDB
func EnsureIndexes(client *mongo.Client) {
	// Define a context with a timeout for index creation operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the collection
	collection := client.Database("slotmachine").Collection("players")

	// Define indexes
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "status", Value: 1}},
		},
	}

	// Create indexes
	for _, index := range indexes {
		_, err := collection.Indexes().CreateOne(ctx, index)
		if err != nil {
			log.Printf("Failed to create index: %v", err)
		}
	}

	log.Println("Indexes ensured in MongoDB")
}
