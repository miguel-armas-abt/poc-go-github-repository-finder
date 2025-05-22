package parameters

import (
	"context"
	"errors"
	"time"

	"com.demo.poc/cmd/parameters/repository/parameters/document"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoParameterRepositoryImpl struct {
	collection *mongo.Collection
}

func NewRepoParameterRepositoryImpl(db *mongo.Database) RepoParameterRepository {
	col := db.Collection("repoParameters")
	return &repoParameterRepositoryImpl{collection: col}
}

func (repository *repoParameterRepositoryImpl) Insert(ctx context.Context, param *document.RepoParameterDocument) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := repository.collection.InsertOne(ctx, param)
	return err
}

func (r *repoParameterRepositoryImpl) FindByRepository(ctx context.Context, repo string) (*document.RepoParameterDocument, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"repository": repo}
	result := r.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, result.Err()
	}
	var param document.RepoParameterDocument
	if err := result.Decode(&param); err != nil {
		return nil, err
	}
	return &param, nil
}

func (r *repoParameterRepositoryImpl) FindByOwner(ctx context.Context, owner string) ([]*document.RepoParameterDocument, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"owner": owner}
	cursor, err := r.collection.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*document.RepoParameterDocument
	for cursor.Next(ctx) {
		var param document.RepoParameterDocument
		if err := cursor.Decode(&param); err != nil {
			return nil, err
		}
		results = append(results, &param)
	}
	return results, nil
}
