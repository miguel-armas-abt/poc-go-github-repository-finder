package document

type RepoMetadataDocument struct {
	RepositoryName string `bson:"repositoryName"`
	Profile        string `bson:"profile"`
	ImageUrl       string `bson:"imageUrl"`
	Label          string `bson:"label"`
	Priority       int    `bson:"priority"`
}
