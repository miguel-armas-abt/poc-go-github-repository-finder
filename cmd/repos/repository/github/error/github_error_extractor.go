package error

import (
	"encoding/json"

	"com.demo.poc/commons/constants"
)

type GitHubErrorExtractor struct{}

const GITHUB_ERROR_WRAPPER_TYPE string = "GIT_HUB_ERROR"

func (e GitHubErrorExtractor) Extract(jsonBody string) (string, string, bool) {
	var resp struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	if err := json.Unmarshal([]byte(jsonBody), &resp); err != nil {
		return constants.EMPTY, constants.EMPTY, false
	}
	if resp.Status == constants.EMPTY && resp.Message == constants.EMPTY {
		return constants.EMPTY, constants.EMPTY, false
	}
	return resp.Status, resp.Message, true
}

func (e GitHubErrorExtractor) Supports(wrapperType string) bool {
	return wrapperType == GITHUB_ERROR_WRAPPER_TYPE
}
