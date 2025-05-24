package validations

type RepoFinderParams struct {
	Profile string `mapstructure:"profile" validate:"required"`
	Label   string `mapstructure:"label" validate:"required"`
}
