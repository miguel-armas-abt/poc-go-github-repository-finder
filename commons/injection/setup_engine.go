package injection

import (
	"net/http"

	"com.demo.poc/commons/interceptor/restclient"
	"com.demo.poc/commons/logging"
	"com.demo.poc/commons/properties"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	logging.InitLogger(logrus.InfoLevel)
	properties.Init()
	http.DefaultClient.Transport = restclient.NewRestClientInterceptor(http.DefaultTransport, &properties.Properties)
	return NewEngine()
}
