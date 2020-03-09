package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Client *mongo.Client
	Db     *mongo.Database
	Ctx    context.Context
}
