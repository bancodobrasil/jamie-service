package loaders

import (
	"context"
)

// Loader ...
type Loader interface {
	Load(ctx context.Context, uuid string, version string) (string, error)
	isConfigured() bool
	checkConfig() error
}
