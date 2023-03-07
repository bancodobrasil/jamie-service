package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bancodobrasil/jamie-service/config"
	"github.com/bancodobrasil/jamie-service/dtos"
	"github.com/bancodobrasil/jamie-service/loaders"
	"github.com/lmmfy/goejs/pkg/contract"
	"github.com/lmmfy/goejs/pkg/interpreter/otto"
)

// Menu ...
type Menu interface {
	Get(ctx context.Context, uuid string, version string) (*dtos.Menu, error)
	Process(ctx context.Context, uuid string, version string, payload *dtos.Eval) (string, error)
}

type menu struct {
	cfg            *config.Config
	loadersManager loaders.Manager
	cacheService   Cache
}

// NewMenu ...
func NewMenu(cfg *config.Config, loadersManager loaders.Manager, cacheService Cache) Menu {
	return &menu{cfg: cfg, loadersManager: loadersManager, cacheService: cacheService}
}

// load ...
func (s *menu) load(ctx context.Context, source string, uuid string, version string) (*dtos.Menu, error) {

	loader, err := s.loadersManager.GetLoader(ctx, source)

	if err != nil {
		return nil, err
	}

	if loader == nil {
		return nil, fmt.Errorf("not configured loader: %s", source)
	}

	return (*loader).Load(ctx, uuid, version)
}

// Get ...
func (s *menu) Get(ctx context.Context, uuid string, version string) (*dtos.Menu, error) {

	if version == "" {
		version = CurrentMenuVersion
	}

	dto, err := s.cacheService.Get(ctx, uuid, version)
	if err != nil {
		return nil, err
	}

	if dto == nil {
		dto, err = s.load(ctx, s.cfg.LoaderSource, uuid, version)
		if err != nil {
			return nil, err
		}

		ttl := time.Duration(s.cfg.Cache.ClosedTTL) * time.Second

		if version == CurrentMenuVersion {
			ttl = time.Duration(s.cfg.Cache.CurrentTTL) * time.Second
		}

		err = s.cacheService.Put(ctx, uuid, version, dto, ttl)
		if err != nil {
			return nil, err
		}
	}

	return dto, nil
}

func (s *menu) Process(ctx context.Context, uuid string, version string, dto *dtos.Eval) (string, error) {
	menu, err := s.Get(ctx, uuid, version)

	if err != nil {
		return "", err
	}

	byteArray, err := json.Marshal(menu)
	if err != nil {
		return "", err
	}

	payload := make(map[string]interface{})

	err = json.Unmarshal(byteArray, &payload)
	if err != nil {
		return "", err
	}

	e := otto.NewDefaultOttoEngine()
	got, err := e.Exec(menu.Template, map[string]interface{}{"menu": payload}, &contract.Option{
		Debug: true,
	})
	if err != nil {
		return "", err
	}

	return got, nil
}
