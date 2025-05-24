package parameters

import (
	"context"
	"errors"
	"time"

	profileErrors "com.demo.poc/cmd/profile/errors"
	"com.demo.poc/cmd/profile/repository/profile/document"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type profileRepositoryImpl struct {
	collection *mongo.Collection
}

func NewProfileRepositoryImpl(db *mongo.Database) ProfileRepository {
	col := db.Collection("profiles")
	return &profileRepositoryImpl{collection: col}
}

func (repository *profileRepositoryImpl) Insert(ctx context.Context, metadata *document.ProfileDocument) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, metadata)
	return err
}

func (repository *profileRepositoryImpl) FindByUsername(ctx context.Context, username string) (*document.ProfileDocument, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var result document.ProfileDocument

	err := repository.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, profileErrors.NewProfileNotFoundError("profile not found")
		}
		return nil, err
	}
	return &result, nil
}
