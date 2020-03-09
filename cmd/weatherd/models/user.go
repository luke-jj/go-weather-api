package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty"`
	FirstName   string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	DateCreated string             `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	LastLogin   string             `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
}

type UserAuthValidation struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
