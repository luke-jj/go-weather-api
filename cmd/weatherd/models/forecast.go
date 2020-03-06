package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Forecast struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"omitempty"`
	Title string             `json:"title,omitempty" bson:"title,omitempty" validate:"gte=2,lte=255"`
	Text  string             `json:"text,omitempty" bson:"text,omitempty" validate:"gte=2,lte=65,535`
}
