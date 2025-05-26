package service

import (
	"context"

	"poc/cmd/profile/dto/request"
	"poc/cmd/profile/dto/response"
)

type ProfileService interface {
	InsertProfile(
		ctx context.Context,
		headers map[string]string,
		insertRequest request.ProfileInsertRequest) error

	FindByUsername(
		ctx context.Context,
		headers map[string]string,
		username string) (*response.ProfileResponse, error)
}
