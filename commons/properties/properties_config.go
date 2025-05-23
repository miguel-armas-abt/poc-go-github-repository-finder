package properties

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Properties ApplicationProperties

func Init(yamlBytes []byte) {
	reader := viper.New()
	reader.SetConfigType("yaml")

	if err := reader.ReadConfig(bytes.NewBuffer(yamlBytes)); err != nil {
		panic(fmt.Sprintf("Error reading embedded application.yaml: %v", err))
	}

	reader.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	reader.AutomaticEnv()

	environmentVariables := []struct{ key, env string }{
		{"server.port", "PORT"},
		{"mongodb.uri", "MONGODB_URI"},
		{"mongodb.database", "MONGODB_DATABASE"},
	}
	for _, variable := range environmentVariables {
		reader.BindEnv(variable.key, variable.env)
	}

	if err := reader.Unmarshal(&Properties); err != nil {
		panic(fmt.Sprintf("Error deserializing config: %v", err))
	}
}
