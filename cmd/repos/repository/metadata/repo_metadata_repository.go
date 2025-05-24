package parameters

import (
	"context"

	params "com.demo.poc/cmd/repos/params"
	"com.demo.poc/cmd/repos/repository/metadata/document"
)

type RepoMetadataRepository interface {
	Insert(ctx context.Context, metadata *document.RepoMetadataDocument) error

	FindByProfileAndLabel(ctx context.Context, params *params.RepoFinderParams) ([]*document.RepoMetadataDocument, error)
}
