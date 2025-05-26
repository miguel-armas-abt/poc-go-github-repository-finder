package config

import (
	"context"
	"fmt"
	"time"

	"poc/commons/custom/properties"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(properties.Properties.Mongo.Uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}
	return client
}
