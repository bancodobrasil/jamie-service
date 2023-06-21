package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bancodobrasil/jamie-service/dtos"
	payloads "github.com/bancodobrasil/jamie-service/payloads/v1"
	"github.com/bancodobrasil/jamie-service/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Menu ...
type Menu interface {
	GetHandler() gin.HandlerFunc
	EvalHandler() gin.HandlerFunc
}

type menu struct {
	service services.Menu
}

// NewMenu ...
func NewMenu(service services.Menu) Menu {
	return &menu{service: service}
}

// LoadMutex ...
// var loadMutex sync.Mutex

// GetHandler godoc
// @Summary 		Evaluate the rulesheet
// @Description 	Receive the params to execute the rulesheet
// @Tags 			eval
// @Accept  		json
// @Produce  		json
// @Param			uuid path string false "uuid"
// @Param 			version path string false "version"
// @Success 		200 {string} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Security 		Authentication Api Key
// @Router 			/menus/{uuid}/{version} [get]
// @Router 			/menus/{uuid} [get]
func (ctrl *menu) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		uuid := c.Param("uuid")

		version := c.Param("version")
		if version == "" {
			version = services.CurrentMenuVersion
		}

		log.Debugf("Get with %s %s\n", uuid, version)

		menu, err := ctrl.service.Get(ctx, uuid, version)
		if err != nil {
			log.Errorf("Erro on load: %v", err)
			c.String(http.StatusInternalServerError, "Error on load menu and/or version")
			return
		}

		c.String(200, menu)

	}

}

// GetHandler godoc
// @Summary 		Evaluate the menu
// @Description 	Receive the params to execute the menu
// @Tags 			eval
// @Accept  		json
// @Produce  		json
// @Param			uuid path string false "uuid"
// @Param 			version path string false "version"
// @Param  			payload body object true "Payload"
// @Success 		200 {string} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Security 		Authentication Api Key
// @Router 			/menus/{uuid}/{version}/eval [post]
// @Router 			/menus/{uuid}/eval [post]
func (ctrl *menu) EvalHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		uuid := c.Param("uuid")

		version := c.Param("version")
		if version == "" {
			version = services.CurrentMenuVersion
		}

		log.Debugf("Eval with %s %s\n", uuid, version)

		decoder := json.NewDecoder(c.Request.Body)
		var t payloads.Eval
		err := decoder.Decode(&t)
		if err != nil {
			log.Errorf("Erro on json decode: %v", err)
			c.Status(http.StatusInternalServerError)
			fmt.Fprint(c.Writer, "Error on json decode")
			return
		}

		dto := dtos.NewProcess(t)

		menu, err := ctrl.service.Process(ctx, uuid, version, &dto)
		if err != nil {
			log.Errorf("Erro on load: %v", err)
			c.String(http.StatusInternalServerError, "Error on load menu and/or version")
			return
		}

		c.String(200, menu)

	}
}
