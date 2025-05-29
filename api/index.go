package handler

import (
	_ "embed"
	"net/http"

	"poc/commons/custom/injection"
)

//go:embed application.yaml
var applicationYAML []byte

// serverless to deploy in Vercel
func Handler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	engine := injection.NewEngine(applicationYAML)

	engine.ServeHTTP(httpResponse, httpRequest)
}
