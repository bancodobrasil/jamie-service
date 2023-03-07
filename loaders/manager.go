package loaders

import (
	"context"
	"fmt"

	"github.com/bancodobrasil/jamie-service/config"
	log "github.com/sirupsen/logrus"
)

// Manager ...
type Manager interface {
	GetLoaders(ctx context.Context) (map[string]Loader, error)
	GetLoader(ctx context.Context, name string) (*Loader, error)
	Register(name string, loader Loader) error
}

type manager struct {
	cfg     *config.Config
	loaders map[string]Loader
}

// NewManager ...
func NewManager(cfg *config.Config) (Manager, error) {
	m := &manager{cfg: cfg}
	err := m.setupLoaders()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *manager) setupLoaders() error {
	m.loaders = make(map[string]Loader)

	var err error

	err = registerFileSystemLoader(m)
	if err != nil {
		return err
	}

	err = registerS3Loader(m)
	if err != nil {
		return err
	}

	if len(m.loaders) == 0 {
		return fmt.Errorf("there aren't configured loaders")
	}

	return nil
}

func (m *manager) GetLoaders(ctx context.Context) (map[string]Loader, error) {
	if m.loaders == nil {
		err := m.setupLoaders()
		if err != nil {
			return nil, err
		}
	}
	return m.loaders, nil
}

func (m *manager) GetLoader(ctx context.Context, name string) (*Loader, error) {
	lMap, err := m.GetLoaders(ctx)
	if err != nil {
		return nil, err
	}
	l, ok := lMap[name]
	if !ok {
		return nil, fmt.Errorf("the loader '%s' isn't configured", name)
	}
	return &l, nil
}

func (m *manager) Register(name string, loader Loader) error {

	if m.loaders == nil {
		return fmt.Errorf("the loaders didn't setup yet")
	}

	m.loaders[name] = loader

	log.Debugf("Loader registred: %s", name)

	return nil
}
