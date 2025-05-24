package parameters

import (
	"context"

	"com.demo.poc/cmd/profile/repository/profile/document"
)

type ProfileRepository interface {
	Insert(ctx context.Context, document *document.ProfileDocument) error

	FindByUsername(ctx context.Context, username string) (*document.ProfileDocument, error)
}
