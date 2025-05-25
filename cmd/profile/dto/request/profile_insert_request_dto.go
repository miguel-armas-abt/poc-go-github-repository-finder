package request

type ProfileInsertRequest struct {
	Username    string       `json:"username" validate:"required"`
	FullName    string       `json:"fullName" validate:"required"`
	CvName      string       `json:"cvName" validate:"required"`
	LinkedinUrl string       `json:"linkedinUrl" validate:"required"`
	RepoFilters []RepoFilter `json:"repoFilters" validate:"required"`
}

type RepoFilter struct {
	Key         string `json:"key" validate:"required"`
	Description string `json:"description" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
}
