package parameters

import (
	"context"

	"com.demo.poc/cmd/parameters/repository/parameters/document"
)

type RepoParameterRepository interface {
	Insert(ctx context.Context, param *document.RepoParameterDocument) error

	FindByRepository(ctx context.Context, repo string) (*document.RepoParameterDocument, error)

	FindByOwner(ctx context.Context, owner string) ([]*document.RepoParameterDocument, error)
}
