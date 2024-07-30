package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"slot-machine-api/models"
)

type PlayerRepository struct {
	client *mongo.Client
}

func NewPlayerRepository(client *mongo.Client) *PlayerRepository {
	return &PlayerRepository{client}
}

func (r *PlayerRepository) CreatePlayer(player *models.Player) error {
	collection := r.client.Database("slotmachine").Collection("players")
	player.ID = primitive.NewObjectID()
	player.Credits = 100
	player.Status = "active"
	_, err := collection.InsertOne(context.Background(), player)
	return err
}

func (r *PlayerRepository) GetPlayerByID(id primitive.ObjectID) (*models.Player, error) {
	var player models.Player
	collection := r.client.Database("slotmachine").Collection("players")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&player)
	return &player, err
}

func (r *PlayerRepository) UpdatePlayerCredits(id primitive.ObjectID, credits int) error {
	collection := r.client.Database("slotmachine").Collection("players")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"credits": credits}})
	return err
}

func (r *PlayerRepository) SuspendPlayer(id primitive.ObjectID) error {
	collection := r.client.Database("slotmachine").Collection("players")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "suspended"}})
	return err
}

func (r *PlayerRepository) ActivatePlayer(id primitive.ObjectID) error {
	collection := r.client.Database("slotmachine").Collection("players")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "active"}})
	return err
}

func InitMongoDB(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
