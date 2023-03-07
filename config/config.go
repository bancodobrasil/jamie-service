package config

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	LoaderSource string `mapstructure:"JAMIE_SERVICE_LOADER_SOURCE"`

	LoadURL          string `mapstructure:"JAMIE_SERVICE_LOAD_URL"`
	LoaderHeaders    http.Header
	LoaderHeadersStr string `mapstructure:"JAMIE_SERVICE_LOADER_HEADERS"`

	Port             string `mapstructure:"PORT"`
	DisableSSLVerify bool   `mapstructure:"JAMIE_DISABLE_SSL_VERIFY"`

	ExternalHost string `mapstructure:"EXTERNAL_HOST"`

	AuthAPIKey string `mapstructure:"JAMIE_SERVICE_API_KEY"`

	Cache *Cache

	FileSystem *FileSystemLoader

	S3 *S3Loader
}

// Cache ...
type Cache struct {
	ClosedTTL  int `mapstructure:"JAMIE_SERVICE_CACHE_CLOSED_TTL"`
	CurrentTTL int `mapstructure:"JAMIE_SERVICE_CACHE_CURRENT_TTL"`
}

// FileSystemLoader ...
type FileSystemLoader struct {
	Path string `mapstructure:"JAMIE_SERVICE_LOADER_FILE_SYSTEM_PATH"`
}

// S3Loader ...
type S3Loader struct {
	Endpoint  string `mapstructure:"JAMIE_SERVICE_LOADER_S3_ENDPOINT"`
	Bucket    string `mapstructure:"JAMIE_SERVICE_LOADER_S3_BUCKET"`
	AccessKey string `mapstructure:"JAMIE_SERVICE_LOADER_S3_ACCESS_KEY"`
	SecretKey string `mapstructure:"JAMIE_SERVICE_LOADER_S3_SECRET_KEY"`
	Secure    bool   `mapstructure:"JAMIE_SERVICE_LOADER_S3_SSL"`
}

var config = &Config{
	Cache:      &Cache{},
	FileSystem: &FileSystemLoader{},
	S3:         &S3Loader{},
}

var loaded = false

// LoadConfig ...
func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("JAMIE_SERVICE_LOADER_SOURCE", "file-system")
	viper.SetDefault("JAMIE_SERVICE_LOADER_FILE_SYSTEM_PATH", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_S3_ENDPOINT", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_S3_BUCKET", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_S3_ACCESS_KEY", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_S3_SECRET_KEY", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_S3_SSL", "true")
	viper.SetDefault("JAMIE_SERVICE_CACHE_CLOSED_TTL", "3600")
	viper.SetDefault("JAMIE_SERVICE_CACHE_CURRENT_TTL", "60")
	viper.SetDefault("JAMIE_SERVICE_LOADER_URL", "")
	viper.SetDefault("JAMIE_SERVICE_LOADER_HEADERS", "")
	viper.SetDefault("JAMIE_SERVICE_DEFAULT_RULES", "")
	viper.SetDefault("PORT", "8005")
	viper.SetDefault("JAMIE_DISABLE_SSL_VERIFY", false)
	viper.SetDefault("EXTERNAL_HOST", "localhost:8005")
	viper.SetDefault("JAMIE_SERVICE_API_KEY", "")

	err = viper.ReadInConfig()
	if err != nil {
		if err2, ok := err.(*os.PathError); !ok {
			err = err2
			log.Errorf("Error on Load Config: %v", err)
			return
		}
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Error on Unmarshal Config: %v", err)
		return
	}

	err = viper.Unmarshal(config.Cache)
	if err != nil {
		log.Fatalf("Error on Unmarshal Config Cache: %v", err)
		return
	}

	err = viper.Unmarshal(config.FileSystem)
	if err != nil {
		log.Fatalf("Error on Unmarshal Config FileSystem: %v", err)
		return
	}

	err = viper.Unmarshal(config.S3)
	if err != nil {
		log.Fatalf("Error on Unmarshal Config S3: %v", err)
		return
	}

	config.LoaderHeaders = make(http.Header)
	resourceLoaderHeaders := strings.Split(config.LoaderHeadersStr, ",")
	for _, value := range resourceLoaderHeaders {
		entries := strings.Split(value, ":")
		if len(entries) == 2 {
			config.LoaderHeaders.Set(entries[0], entries[1])
		}
	}

	return
}

// GetConfig ...
func GetConfig() *Config {
	if !loaded {
		err := LoadConfig()
		loaded = true
		if err != nil {
			panic(fmt.Sprintf("load config error: %s", err))
		}
	}
	return config
}
