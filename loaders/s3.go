package loaders

import (
	"context"
	"fmt"
	"io"

	"github.com/bancodobrasil/jamie-service/config"
	log "github.com/sirupsen/logrus"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// S3Loader ...
const S3Loader = "s3"

// S3 ...
type S3 interface {
	Loader
}

type s3 struct {
	cfg config.S3Loader
}

// NewS3 ...
func NewS3(cfg config.S3Loader) (S3, error) {

	l := &s3{cfg: cfg}

	if !l.isConfigured() {
		return nil, nil
	}

	err := l.checkConfig()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *s3) isConfigured() bool {
	return l.cfg.Endpoint != "" &&
		l.cfg.Bucket != "" &&
		l.cfg.AccessKey != "" &&
		l.cfg.SecretKey != ""
}

func (l *s3) checkConfig() error {
	// TODO check bucket exists
	return nil
}

func (l *s3) Load(ctx context.Context, uuid string, version string) (string, error) {
	log.Debugf("Loading from %s > uuid:%s version: %s", S3Loader, uuid, version)

	log.Tracef("S3> endpoint: %s access_key: %s bucket: %s object: %s", l.cfg.Endpoint, l.cfg.AccessKey, l.cfg.Bucket, fmt.Sprintf("%s/%s.json", uuid, version))

	// Initialize client object.
	client, err := minio.New(l.cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(l.cfg.AccessKey, l.cfg.SecretKey, ""),
		Secure: l.cfg.Secure,
	})
	if err != nil {
		log.Errorf("s3: error on new client: %s", err)
		return "", err
	}

	obj, err := client.GetObject(ctx, l.cfg.Bucket, fmt.Sprintf("%s/%s.jamie", uuid, version), minio.GetObjectOptions{})
	if err != nil {
		log.Errorf("s3: error on get: %s", err)
		return "", err
	}
	defer obj.Close()

	byteValue, err := io.ReadAll(obj)
	if err != nil {
		log.Errorf("s3: error on read: %s", err)
		return "", err
	}

	return string(byteValue), nil
}

func registerS3Loader(m *manager) error {
	loader, err := NewS3(*m.cfg.S3)
	if err != nil {
		return err
	}
	if loader != nil {
		m.Register(S3Loader, loader)
	}
	return nil
}
