package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	errorWrapper "com.demo.poc/cmd/repos/repository/github/error"
	"com.demo.poc/cmd/repos/repository/github/wrapper/response"
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

func (repository *gitHubRepositoryImpl) FindRepositoriesByProfile(
	ctx context.Context,
	headers map[string]string,
	profile string) ([]response.RepoResponseWrapper, error) {

	restClient, err := properties.Properties.SearchRestClient(SERVICE_NAME)
	if err != nil {
		return nil, err
	}

	perPage := restClient.Request.Params["per_page"]

	var result []response.RepoResponseWrapper
	response, err := repository.restyClient.
		SetBaseURL(restClient.Request.Endpoint).
		R().
		SetHeaders(filler.FillHeaders(headers, restClient)).
		SetResult(&result).
		Get(fmt.Sprintf("/%s/repos?per_page=%s", profile, perPage))

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, repository.errorHandler.HandleError(
			response,
			SERVICE_NAME,
			errorWrapper.GITHUB_ERROR_WRAPPER_TYPE,
		)
	}

	return result, nil
}
