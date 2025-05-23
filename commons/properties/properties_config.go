package properties

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Properties ApplicationProperties

func Init() {
	reader := viper.New()
	reader.SetConfigName("application")
	reader.SetConfigType("yaml")
	reader.AddConfigPath(".")
	reader.AddConfigPath("./resources")
	reader.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	reader.AutomaticEnv()

	reader.BindEnv("server.port", "PORT")
	reader.BindEnv("mongodb.uri", "MONGODB_URI")
	reader.BindEnv("mongodb.database", "MONGODB_DATABASE")

	if err := reader.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading application.yaml: %v", err))
	}

	if err := reader.Unmarshal(&Properties); err != nil {
		panic(fmt.Sprintf("Error deserializing config: %v", err))
	}
}
