package config

import (
	"context"
	"fmt"
	"time"

	properties "com.demo.poc/commons/properties"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(props *properties.ApplicationProperties) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(props.Mongo.Uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}
	return client
}
