package v1

import (
	log "github.com/sirupsen/logrus"

	"github.com/bancodobrasil/jamie-service/config"
	v1 "github.com/bancodobrasil/jamie-service/controllers/v1"
	"github.com/bancodobrasil/jamie-service/loaders"
	"github.com/bancodobrasil/jamie-service/services"
	"github.com/gin-gonic/gin"
)

func evalRouter(router *gin.RouterGroup) {

	cfg := config.GetConfig()
	loadersManager, err := loaders.NewManager(cfg)

	if err != nil {
		log.Fatal("error on init loader manager: ", err)
	}

	cacheService := services.NewCache()

	service := services.NewMenu(cfg, loadersManager, cacheService)
	controller := v1.NewMenu(service)

	router.GET("/:uuid/:version", controller.GetHandler())
	router.GET("/:uuid/:version/", controller.GetHandler())
	router.GET("/:uuid", controller.GetHandler())
	router.GET("/:uuid/", controller.GetHandler())

	router.POST("/:uuid/:version/eval", controller.EvalHandler())
	router.POST("/:uuid/:version/eval/", controller.EvalHandler())
	router.POST("/:uuid/eval", controller.EvalHandler())
	router.POST("/:uuid/eval/", controller.EvalHandler())

}
