package parameters

import (
	"context"

	"com.demo.poc/cmd/parameters/repository/parameters/document"
	params "com.demo.poc/cmd/repos/params"
)

type ParameterRepository interface {
	Insert(ctx context.Context, param *document.ParameterDocument) error

	FindByOwnerAndLabel(ctx context.Context, params *params.RepoFinderParams) ([]*document.ParameterDocument, error)
}
