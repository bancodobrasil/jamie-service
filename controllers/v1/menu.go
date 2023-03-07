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
// @Router 			/{uuid}/{version} [get]
// @Router 			/{uuid} [get]
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

		c.JSON(200, menu)

		//loadMutex.Lock()

		// menu := services.EvalService.GetKnowledgeLibrary().GetMenu(uuid, version)

		// if !(len(menu.RuleEntries) > 0) {

		// 	err := services.EvalService.LoadRemoteGRL(uuid, version)
		// 	if err != nil {
		// 		log.Errorf("Erro on load: %v", err)
		// 		c.String(http.StatusInternalServerError, "Error on load menu and/or version")
		// 		loadMutex.Unlock()
		// 		return
		// 	}

		// 	menu = services.EvalService.GetKnowledgeLibrary().GetMenu(uuid, version)

		// 	if !(len(menu.RuleEntries) > 0) {
		// 		c.Status(http.StatusNotFound)
		// 		fmt.Fprint(c.Writer, "Menu or version not founded!")
		// 		loadMutex.Unlock()
		// 		return
		// 	}
		// }

		// loadMutex.Unlock()

		// decoder := json.NewDecoder(c.Request.Body)
		// var t payloads.Eval
		// err := decoder.Decode(&t)
		// if err != nil {
		// 	log.Errorf("Erro on json decode: %v", err)
		// 	c.Status(http.StatusInternalServerError)
		// 	fmt.Fprint(c.Writer, "Error on json decode")
		// 	return
		// }
		// log.Debugln(t)

		// ctx := types.NewContextFromMap(t)
		// ctx.RawContext = c.Request.Context()

		// result, err := services.EvalService.Eval(ctx, menu)
		// if err != nil {

		// 	log.Errorf("Error on eval: %v", err)
		// 	c.Status(http.StatusInternalServerError)
		// 	fmt.Fprint(c.Writer, "Error on eval")
		// 	return
		// }

		// log.Debug("Context:\n\t", ctx.GetEntries(), "\n\n")
		// log.Debug("Features:\n\t", result.GetFeatures(), "\n\n")

		// responseCode := http.StatusOK

		// if result.Has("requiredParamErrors") {
		// 	responseCode = http.StatusBadRequest
		// }

		// c.JSON(responseCode, result.GetFeatures())

		//c.String(200, "OK")
	}

}

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
// @Router 			/{uuid}/{version} [get]
// @Router 			/{uuid} [get]
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

		dto := dtos.NewEval(t)

		menu, err := ctrl.service.Process(ctx, uuid, version, &dto)
		if err != nil {
			log.Errorf("Erro on load: %v", err)
			c.String(http.StatusInternalServerError, "Error on load menu and/or version")
			return
		}

		c.JSON(200, menu)

	}
}
