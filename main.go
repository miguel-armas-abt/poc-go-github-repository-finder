package main

import (
	"strings"

	"com.demo.poc/commons/constants"
	"com.demo.poc/commons/injection"
	properties "com.demo.poc/commons/properties"
)

var engine = injection.InitEngine()

func main() {
	router := engine

	serverPort := properties.Properties.Server.Port
	if !strings.HasPrefix(serverPort, constants.COLON) {
		serverPort = constants.COLON + serverPort
	}
	router.Run(serverPort)
}
