package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Forecast struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty" validate:"required,gte=2,lte=255"`
	Location     string             `json:"location,omitempty" bson:"location,omitempty" validate:"gte=2,lte=255"`
	Text         string             `json:"text,omitempty" bson:"text,omitempty" validate:"gte=2,lte=65535"`
	CreatedBy    User               `json:"createdBy,omitempty", bson"createdBy,omitempty" validate:"omitempty"`
	DateCreated  string             `json:"dateCreated,omitempty" bson:"dateCreated,omitempty" validate:"omitempty"`
	DateModified string             `json:"dateModified,omitempty" bson:"dateModified,omitempty" validate:"omitempty"`
}
