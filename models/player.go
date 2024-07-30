package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Credits int                `json:"credits" bson:"credits"`
	Status  string             `json:"status" bson:"status"`
}

type PlayResult struct {
	Result string `json:"result"`
	Payout int    `json:"payout"`
}
