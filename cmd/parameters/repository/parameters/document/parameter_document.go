package document

type ParameterDocument struct {
	RepositoryName string `bson:"repositoryName"`
	Owner          string `bson:"owner"`
	ImageUrl       string `bson:"imageUrl"`
	Label          string `bson:"label"`
}
