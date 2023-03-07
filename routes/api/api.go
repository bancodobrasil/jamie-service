package api

import (
	v1 "github.com/bancodobrasil/jamie-service/routes/api/v1"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(router *gin.RouterGroup) {
	v1.Router(router.Group("/v1"))
}
