package response

type ProfileResponse struct {
	FullName    string       `json:"fullName"`
	CvUrl       string       `json:"cvUrl"`
	RepoFilters []RepoFilter `json:"repoFilters"`
}

type RepoFilter struct {
	Key         string `json:"key" validate:"required"`
	Description string `json:"description" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
}
