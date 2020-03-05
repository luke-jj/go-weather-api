package models

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Forecast struct {
	ID    primitive.ObjectID `json:"id",omitempty`
	Title string             `json:"title"`
	Text  string             `json:"text"`
}
