package loaders

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/bancodobrasil/jamie-service/config"
	"github.com/bancodobrasil/jamie-service/dtos"
	log "github.com/sirupsen/logrus"
)

// FileSystemLoader ...
const FileSystemLoader = "file-system"

// FileSystem ...
type FileSystem interface {
	Loader
}

type fileSystem struct {
	cfg config.FileSystemLoader
}

// NewFileSystem ...
func NewFileSystem(cfg config.FileSystemLoader) (FileSystem, error) {

	l := &fileSystem{
		cfg: cfg,
	}

	if !l.isConfigured() {
		return nil, nil
	}

	err := l.checkConfig()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *fileSystem) isConfigured() bool {
	return l.cfg.Path != ""
}

func (l *fileSystem) checkConfig() error {
	if _, err := os.Stat(l.cfg.Path); errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (l *fileSystem) Load(ctx context.Context, uuid string, version string) (*dtos.Menu, error) {
	log.Debugf("Loading from %s > uuid:%s version: %s", FileSystemLoader, uuid, version)

	filePath := path.Join(l.cfg.Path, uuid, version+".jamie")

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("this version not exists: %s", filePath)
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	dto := &dtos.Menu{}

	err = json.Unmarshal(byteValue, dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func registerFileSystemLoader(m *manager) error {
	fsLoader, err := NewFileSystem(*m.cfg.FileSystem)
	if err != nil {
		return err
	}
	if fsLoader != nil {
		m.Register(FileSystemLoader, fsLoader)
	}
	return nil
}
