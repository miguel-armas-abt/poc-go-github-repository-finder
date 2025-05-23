package main

import (
	"net/http"

	"com.demo.poc/commons/injection"
)

var engine = injection.InitEngine()

// serverless to deploy in Vercel
func Handler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	engine.ServeHTTP(httpResponse, httpRequest)
}
