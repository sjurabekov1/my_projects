package config

import (
	"fmt"
	"os"
	"time"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
	"hyssa/hyssa_go_billing_service/pkg/logger/options"
)

type AppMode string
type AuthMode string

const (
	DEVELOPMENT AppMode = "DEVELOPMENT"
	PRODUCTION  AppMode = "PRODUCTION"

	ACCESS   AuthMode = "access"
	REFRESH  AuthMode = "refresh"
	PASSCODE AuthMode = "passcode"
)

var (
	TimeoutDuration time.Duration = 7 * time.Second
	CacheCount      int           = 100_000
	TimeToLiveMins                = time.Duration(5) * time.Minute
)

type Config struct {
	Logging options.Logging `yaml:"logging"`

	Project struct {
		Name           string        `env:"PROJECT_NAME" yaml:"name"`
		Mode           string        `env:"APPLICATION_MODE,default=PRODUCTION"`
		Version        string        `env:"APPLICATION_VERSION" yaml:"version"`
		Timout         time.Duration `env:"TIME_OUT" yaml:"timeout"`
		SwaggerEnabled bool          `env:"SWAGGER_ENABLED" yaml:"swagger_enabled"`
		MaxBodyMB      int           `yaml:"max_body_mb"`
		CacheTimeout   int           `yaml:"cache_timeout"`
	} `yaml:"project"`

	Http struct {
		HTTP_HOST string `env:"HTTP_HOST" yaml:"host"`
		HTTP_PORT int    `env:"HTTP_PORT" yaml:"port"`

		URL string `env:"HTTP_URL" yaml:"url"`
	} `yaml:"http"`

	GRPC struct {
		GRPC_HOST string `env:"GRPC_HOST" yaml:"host"`
		GRPC_PORT int    `env:"GRPC_PORT" yaml:"port"`

		URL string `env:"GRPC_URL" yaml:"url"`
	} `yaml:"grpc"`

	PSQL struct {
		URI string `env:"PSQL_URI"`
	}

	AuthConfig struct {
		AccessSecretKey       string `env:"ACCESS_SECRET_KEY,default=access_key"`
		AccessKeyExpireMins   int    `env:"ACCESS_KEY_EXPIRE_MINS,default=180"`
		AccessKeyExpireDays   int    `env:"ACCESS_KEY_EXPIRE_DAYS,default=3"`
		RefreshSecretKey      string `env:"REFRESH_SECRET_KEY,default=refresh_key"`
		RefreshKeyExpireDays  int    `env:"REFRESH_KEY_EXPIRE,default=7"`
		PasscodeSecretKey     string `env:"PASSCODE_SECRET_KEY,default=passcode_key"`
		PasscodeKeyExpireMins int    `env:"PASSCODE_KEY_EXPIRE,default=3"`
	}
}

func Load() *Config {
	var cfg Config

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		logger.Log.Warn(".env file is not found")
	}

	appMode := getAppMode()
	configPath, err := getConfigPath(appMode)
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic("unmarshal from environment error")
	}

	if err := validateConfig(&cfg); err != nil {
		panic(err)
	}

	TimeoutDuration = cfg.Project.Timout

	cfg.MakeGrpcURL()
	cfg.MakeHttpURL()

	return &cfg
}

func (c *Config) MakeGrpcURL() {
	c.GRPC.URL = fmt.Sprintf("%s:%d", c.GRPC.GRPC_HOST, c.GRPC.GRPC_PORT)
}

func (c *Config) MakeHttpURL() {
	c.Http.URL = fmt.Sprintf("%s:%d", c.Http.HTTP_HOST, c.Http.HTTP_PORT)
}

func getAppMode() AppMode {
	mode := AppMode(os.Getenv("APPLICATION_MODE"))

	if mode != DEVELOPMENT {
		mode = PRODUCTION
	}

	return mode
}

func getConfigPath(appMode AppMode) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	suffix := "dev"
	if appMode == PRODUCTION {
		suffix = "prod"
	}
	fmt.Printf("%s/configs/app_configs_%s.yaml\n", path, suffix)
	return fmt.Sprintf("%s/configs/app_configs_%s.yaml", path, suffix), nil
}
