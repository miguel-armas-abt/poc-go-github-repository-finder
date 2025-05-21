package repository

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"com.demo.poc/cmd/products/repository/github/wrapper/response"
	"com.demo.poc/commons/properties"
	restClientErrors "com.demo.poc/commons/restclient/errors"
	"com.demo.poc/commons/restclient/filler"
)

const SERVICE_NAME = "github-users"

type gitHubRepositoryImpl struct {
	restyClient  *resty.Client
	properties   *properties.ApplicationProperties
	errorHandler *restClientErrors.RestclientErrorHandler
}

func NewGitHubRepositoryImpl(
	properties *properties.ApplicationProperties,
	errorHandler *restClientErrors.RestclientErrorHandler) GitHubRepository {

	httpClient := http.DefaultClient
	restyClient := resty.New().
		SetTransport(httpClient.Transport)

	return &gitHubRepositoryImpl{
		restyClient:  restyClient,
		properties:   properties,
		errorHandler: errorHandler,
	}
}

func (repository *gitHubRepositoryImpl) FindRepositoriesByOwner(
	headers map[string]string,
	owner string) ([]response.RepoResponseWrapper, error) {

	restClient, err := properties.Properties.SearchRestClient(SERVICE_NAME)
	if err != nil {
		return nil, err
	}

	filledHeaders := filler.FillHeaders(headers, restClient)

	var result []response.RepoResponseWrapper
	response, err := repository.restyClient.
		SetBaseURL(restClient.Request.Endpoint).
		R().
		SetHeaders(filledHeaders).
		SetResult(&result).
		Get(fmt.Sprintf("/%s/repos", owner))

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, repository.errorHandler.HandleError(
			response,
			SERVICE_NAME,
			"GitHubError",
		)
	}

	return result, nil
}
