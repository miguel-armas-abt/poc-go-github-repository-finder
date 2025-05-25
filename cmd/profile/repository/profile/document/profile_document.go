package document

type ProfileDocument struct {
	Username    string       `bson:"username"`
	FullName    string       `bson:"fullName"`
	CvUrl       string       `bson:"cvUrl"`
	GitHubUrl   string       `bson:"gitHubUrl"`
	LinkedinUrl string       `bson:"linkedinUrl"`
	RepoFilters []RepoFilter `bson:"repoFilters"`
}

type RepoFilter struct {
	Key         string `bson:"key"`
	Description string `bson:"description"`
	Summary     string `bson:"summary"`
	Priority    int    `bson:"priority"`
}
