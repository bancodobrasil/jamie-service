package loaders

import (
	"context"

	"github.com/bancodobrasil/jamie-service/dtos"
)

// Loader ...
type Loader interface {
	Load(ctx context.Context, uuid string, version string) (*dtos.Menu, error)
	isConfigured() bool
	checkConfig() error
}
