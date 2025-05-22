package document

type RepoParameterDocument struct {
	Repository string `bson:"repository"`
	Owner      string `bson:"owner"`
	Img        string `bson:"img"`
}
