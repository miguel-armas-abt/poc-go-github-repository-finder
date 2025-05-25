package properties

import (
	"strings"

	"com.demo.poc/commons/constants"
	coreErrors "com.demo.poc/commons/errors/errors"
	logging "com.demo.poc/commons/properties/logging"
	mongoDb "com.demo.poc/commons/properties/mongo"
	"com.demo.poc/commons/properties/restclient"
)

type ApplicationProperties struct {
	Server            ServerProperties                 `mapstructure:"server"`
	ProjectType       ProjectType                      `mapstructure:"projectType"`
	Logging           logging.LoggingTemplate          `mapstructure:"logging"`
	ErrorMessages     map[string]string                `mapstructure:"errorMessages"`
	RestClients       map[string]restclient.RestClient `mapstructure:"restClients"`
	Mongo             mongoDb.Mongo                    `mapstructure:"mongodb"`
	MultimediaStorage string                           `mapstructure:"multimediaStorage"`
	GitHubDomain      string                           `mapstructure:"gitHubDomain"`
}

type ServerProperties struct {
	Port string `mapstructure:"port"`
}

func (properties *ApplicationProperties) IsLoggerEnabled(logType string) bool {
	if properties.Logging.LoggingType == nil {
		return true
	}
	key := strings.ReplaceAll(logType, constants.DOT, constants.MIDDLE_DASH)
	enabled, exists := properties.Logging.LoggingType[key]
	if !exists {
		return true
	}
	return enabled
}

func (properties *ApplicationProperties) SearchRestClient(serviceName string) (*restclient.RestClient, error) {
	restclient, exists := properties.RestClients[serviceName]
	if !exists {
		return nil, coreErrors.NoSuchRestClientError(serviceName)
	}
	return &restclient, nil
}
