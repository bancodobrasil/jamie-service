package main

import (
	"os"

	ginMonitor "github.com/bancodobrasil/gin-monitor"
	"github.com/bancodobrasil/goauth"
	"github.com/bancodobrasil/jamie-service/config"
	_ "github.com/bancodobrasil/jamie-service/docs"
	"github.com/bancodobrasil/jamie-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func setupLog() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	// log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

// @title Jamie Service

// @version 1.0

// @description Service Project provide the menus

// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @contact.url http://www.swagger.io/support

// @contact.email support@swagger.io

// @license.name Apache 2.0

// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000

// @BasePath /api/v1

// @securityDefinitions.apikey Authentication Api Key
// @in header
// @name X-API-Key

// @x-extension-openapi {"example": "value on a json format"}

func main() {

	setupLog()

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
	}

	cfg := config.GetConfig()

	// if cfg.DefaultRules != "" {
	// 	defaultGRL := cfg.DefaultRules
	// 	log.Debugf("Carregando '%s' como folha de regras default!", defaultGRL)
	// 	err := services.EvalService.LoadLocalGRL(defaultGRL, services.DefaultMenuName, services.CurrentMenuVersion)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	log.Warnln("Não foram carregadas regras default!")
	// }

	monitor, err := ginMonitor.New("v0.3.2-rc1", ginMonitor.DefaultErrorMessageKey, ginMonitor.DefaultBuckets)
	if err != nil {
		log.Panic(err)
	}

	gin.DefaultWriter = log.StandardLogger().WriterLevel(log.DebugLevel)
	gin.DefaultErrorWriter = log.StandardLogger().WriterLevel(log.ErrorLevel)

	router := gin.New()

	goauth.BootstrapMiddleware()

	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())
	router.Use(monitor.Prometheus())
	router.GET("metrics", gin.WrapH(promhttp.Handler()))
	routes.SetupRoutes(router)
	routes.APIRoutes(router)

	port := cfg.Port

	router.Run(":" + port)

}
