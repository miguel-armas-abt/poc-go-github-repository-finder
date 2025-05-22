package main

import (
	"log"
	"net/http"
	"strings"

	"com.demo.poc/commons/constants"
	"com.demo.poc/commons/injection"
	"com.demo.poc/commons/interceptor/restclient"
	"com.demo.poc/commons/logging"
	properties "com.demo.poc/commons/properties"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.InitLogger(logrus.InfoLevel)

	if err := properties.Init(); err != nil {
		log.Fatalf("properties load error: %v", err)
	}

	http.DefaultClient.Transport = restclient.NewRestClientInterceptor(http.DefaultTransport, &properties.Properties)

	router := injection.NewEngine()

	serverPort := properties.Properties.Server.Port
	if !strings.HasPrefix(serverPort, constants.COLON) {
		serverPort = constants.COLON + serverPort
	}
	router.Run(serverPort)
}
