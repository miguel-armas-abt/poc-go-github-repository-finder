package parameters

import (
	"context"

	params "poc/cmd/repos/params"
	"poc/cmd/repos/repository/metadata/document"
)

type RepoMetadataRepository interface {
	Insert(ctx context.Context, metadata *document.RepoMetadataDocument) error

	FindByProfileAndLabel(ctx context.Context, params *params.RepoFinderParams) ([]*document.RepoMetadataDocument, error)
}
