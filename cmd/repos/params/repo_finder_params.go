package validations

type RepoFinderParams struct {
	Owner string `mapstructure:"owner" validate:"required"`
	Label string `mapstructure:"label" validate:"required"`
}
