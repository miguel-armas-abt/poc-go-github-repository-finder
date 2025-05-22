package parameters

import (
	"context"
	"time"

	"com.demo.poc/cmd/parameters/repository/parameters/document"
	params "com.demo.poc/cmd/repos/params"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type parameterRepositoryImpl struct {
	collection *mongo.Collection
}

func NewParameterRepositoryImpl(db *mongo.Database) ParameterRepository {
	col := db.Collection("repoParameters")
	return &parameterRepositoryImpl{collection: col}
}

func (repository *parameterRepositoryImpl) Insert(ctx context.Context, param *document.ParameterDocument) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, param)
	return err
}

func (repository *parameterRepositoryImpl) FindByOwnerAndLabel(ctx context.Context, params *params.RepoFinderParams) ([]*document.ParameterDocument, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"owner": params.Owner, "label": params.Label}
	cursor, err := repository.collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*document.ParameterDocument
	for cursor.Next(ctx) {
		var param document.ParameterDocument
		if err := cursor.Decode(&param); err != nil {
			return nil, err
		}
		results = append(results, &param)
	}
	return results, nil
}
