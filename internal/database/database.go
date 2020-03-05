package database

import (
	"context"
	"log"
	"time"

	c "github.com/luke-jj/go-weather-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(config *c.Config) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	config.Client = client
	config.Ctx = ctx
	config.Db = client.Database(config.MONGO_DBNAME)
}
