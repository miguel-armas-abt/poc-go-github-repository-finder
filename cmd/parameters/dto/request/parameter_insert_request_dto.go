package request

type ParameterInsertRequest struct {
	RepositoryName string `json:"repositoryName" validate:"required"`
	Owner          string `json:"owner" validate:"required"`
	Label          string `json:"label" validate:"required"`
}
