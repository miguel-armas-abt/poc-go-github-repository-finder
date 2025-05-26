package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	errorWrapper "poc/cmd/repos/repository/github/error"
	"poc/cmd/repos/repository/github/wrapper/response"
	restClientErrors "poc/commons/core/restclient/errors"
	"poc/commons/core/restclient/filler"
	"poc/commons/custom/properties"
)

const SERVICE_NAME = "github-users"

type gitHubRepositoryImpl struct {
	restyClient  *resty.Client
	errorHandler *restClientErrors.RestclientErrorHandler
}

func NewGitHubRepositoryImpl(
	errorHandler *restClientErrors.RestclientErrorHandler) GitHubRepository {

	httpClient := http.DefaultClient
	restyClient := resty.New().
		SetTransport(httpClient.Transport)

	return &gitHubRepositoryImpl{
		restyClient:  restyClient,
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
