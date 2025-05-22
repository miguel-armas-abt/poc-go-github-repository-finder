package request

type ParameterInsertRequest struct {
	RepositoryName string `json:"repositoryName" validate:"required"`
	Owner          string `json:"owner" validate:"required"`
	ImageUrl       string `json:"imageUrl" validate:"required"`
	Label          string `json:"label" validate:"required"`
}
