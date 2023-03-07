package routes

import (
	"github.com/bancodobrasil/jamie-service/controllers"
	"github.com/gin-gonic/gin"
)

func homeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeHandler())
}
