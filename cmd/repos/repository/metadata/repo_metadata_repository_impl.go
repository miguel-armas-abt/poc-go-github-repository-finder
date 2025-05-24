package parameters

import (
	"context"
	"time"

	params "com.demo.poc/cmd/repos/params"
	"com.demo.poc/cmd/repos/repository/metadata/document"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoMetadataRepositoryImpl struct {
	collection *mongo.Collection
}

func NewRepoMetadataRepositoryImpl(db *mongo.Database) RepoMetadataRepository {
	col := db.Collection("repoMetadata")
	return &repoMetadataRepositoryImpl{collection: col}
}

func (repository *repoMetadataRepositoryImpl) Insert(ctx context.Context, metadata *document.RepoMetadataDocument) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, metadata)
	return err
}

func (repository *repoMetadataRepositoryImpl) FindByProfileAndLabel(ctx context.Context, params *params.RepoFinderParams) ([]*document.RepoMetadataDocument, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"profile": params.Profile, "label": params.Label}
	cursor, err := repository.collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*document.RepoMetadataDocument
	for cursor.Next(ctx) {
		var param document.RepoMetadataDocument
		if err := cursor.Decode(&param); err != nil {
			return nil, err
		}
		results = append(results, &param)
	}
	return results, nil
}
