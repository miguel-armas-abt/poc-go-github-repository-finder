package mongo

type Mongo struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}
