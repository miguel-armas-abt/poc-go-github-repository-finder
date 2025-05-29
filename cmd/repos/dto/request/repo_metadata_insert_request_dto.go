package request

type RepoMetadataInsertRequest struct {
	RepositoryName string `json:"repositoryName" validate:"required"`
	Profile        string `json:"profile" validate:"required"`
	Label          string `json:"label" validate:"required"`
	Priority       int    `json:"priority" validate:"required"`
}
