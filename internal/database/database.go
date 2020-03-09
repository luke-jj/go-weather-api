package database

import (
	"context"
	"log"
	"time"

	c "github.com/luke-jj/go-weather-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(config *c.Config) *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to MongoDB...")

	db := Database{
		Client: client,
		Ctx:    ctx,
		Db:     client.Database(config.MONGO_DBNAME),
	}

	return &db
}
