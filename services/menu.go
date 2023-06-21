package services

import (
	"bytes"
	"context"
	"fmt"
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
	rullerClient := featws.NewRullerClient(cfg.FeatWSRullerURL, cfg.FeatWSRullerAPIKey)

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

	if content == nil {
		content = ""
	}

	return content.(string), nil
}

func (s *menu) Process(ctx context.Context, uuid string, version string, dto *dtos.Eval) (string, error) {
	templateContent, err := s.Get(ctx, uuid, version)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New(uuid).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	// TODO call featws
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
