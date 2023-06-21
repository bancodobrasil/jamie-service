package services

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/bancodobrasil/jamie-service/clients/featws"
	"github.com/bancodobrasil/jamie-service/config"
	"github.com/bancodobrasil/jamie-service/dtos"
	"github.com/bancodobrasil/jamie-service/loaders"
)

// Menu ...
type Menu interface {
	Get(ctx context.Context, uuid string, version string) (string, error)
	Process(ctx context.Context, uuid string, version string, payload *dtos.Eval) (string, error)
}

type menu struct {
	cfg            *config.Config
	loadersManager loaders.Manager
	cacheService   Cache
	rullerClient   *featws.RullerClient
}

// NewMenu ...
func NewMenu(cfg *config.Config, loadersManager loaders.Manager, cacheService Cache) Menu {
	var rullerClient *featws.RullerClient = nil

	if cfg.FeatWSRullerURL != "" {
		featws.NewRullerClient(cfg.FeatWSRullerURL, cfg.FeatWSRullerAPIKey)
	}

	return &menu{cfg: cfg, loadersManager: loadersManager, cacheService: cacheService, rullerClient: rullerClient}
}

// load ...
func (s *menu) load(ctx context.Context, source string, uuid string, version string) (string, error) {

	loader, err := s.loadersManager.GetLoader(ctx, source)

	if err != nil {
		return "", err
	}

	if loader == nil {
		return "", fmt.Errorf("not configured loader: %s", source)
	}

	return (*loader).Load(ctx, uuid, version)
}

// Get ...
func (s *menu) Get(ctx context.Context, uuid string, version string) (string, error) {

	if version == "" {
		version = CurrentMenuVersion
	}

	content, err := s.cacheService.Get(ctx, uuid, version)
	if err != nil {
		return "", err
	}

	if content == nil {
		content, err = s.load(ctx, s.cfg.LoaderSource, uuid, version)
		if err != nil {
			return "", err
		}

		ttl := time.Duration(s.cfg.Cache.ClosedTTL) * time.Second

		if version == CurrentMenuVersion {
			ttl = time.Duration(s.cfg.Cache.CurrentTTL) * time.Second
		}

		err = s.cacheService.Put(ctx, uuid, version, content, ttl)
		if err != nil {
			return "", err
		}
	}

	return content.(string), nil
}

// Process ...
func (s *menu) Process(ctx context.Context, uuid string, version string, dto *dtos.Eval) (string, error) {

	content, err := s.cacheService.Get(ctx, uuid, version)
	if err != nil {
		return "", err
	}

	if content == nil {
		content, err = s.Get(ctx, uuid, version)
		if err != nil {
			return "", err
		}
	}

	var response string

	if s.rullerClient != nil {
		evalResult, err := s.rullerClient.Eval("jamie-menu-"+uuid, version, map[string]interface{}{})
		if err != nil {
			return "", err
		}

		response, err = s.processTemplateConditions(
			uuid,
			content.(string),
			evalResult,
		)
		if err != nil {
			return "", err
		}
	} else {
		response = content.(string)
	}

	return response, nil
}

// processTemplateConditions ...
func (s *menu) processTemplateConditions(uuid string, templateContent string, evalResult *featws.EvalPayload) (string, error) {
	tmpl, err := template.New(uuid).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, evalResult)
	if err != nil {
		return "", err
	}

	result := s.formatJson(buf.String())

	return result, nil
}

// formatJson ...
func (s *menu) formatJson(json string) string {
	if (strings.HasPrefix(json, "{") && strings.HasSuffix(json, "}")) ||
		(strings.HasPrefix(json, "[") && strings.HasSuffix(json, "]")) {
		r := regexp.MustCompile(",(?=\\s*?[\\]}])")
		return r.ReplaceAllString(json, "")
	}
	return json
}
