package config

import (
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mwinyimoha/commons/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName    string `mapstructure:"SERVICE_NAME" validate:"required"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION" validate:"required"`
	Debug          bool   `mapstructure:"DEBUG"`
	AppTimeout     int    `mapstructure:"APP_TIMEOUT" validate:"required,gt=0"`
	ServerPort     int    `mapstructure:"SERVER_PORT" validate:"required,gt=0,lt=65536"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	DatabaseName   string `mapstructure:"DATABASE_NAME"`
	StoreType      string `mapstructure:"STORE_TYPE" validate:"required,oneof=inmemory mongo"`
	APIType        string `mapstructure:"API_TYPE" validate:"required,oneof=rest rpc"`
}

func New(val *validator.Validate) (*Config, error) {
	v := viper.New()
	v.SetConfigType("env")

	v.SetDefault("SERVICE_NAME", "")
	v.SetDefault("SERVICE_VERSION", "0.1.0")
	v.SetDefault("DEBUG", true)
	v.SetDefault("APP_TIMEOUT", 10)
	v.SetDefault("SERVER_PORT", 8080)
	v.SetDefault("DATABASE_URL", "")
	v.SetDefault("DATABASE_NAME", "")
	v.SetDefault("STORE_TYPE", "inmemory")
	v.SetDefault("API_TYPE", "rest")

	v.AutomaticEnv()

	debug := true
	if raw := os.Getenv("DEBUG"); raw != "" {
		val, err := strconv.ParseBool(raw)
		if err == nil {
			debug = val
		}
	}

	if debug {
		configPath := "./"
		v.AddConfigPath(configPath)

		if err := v.ReadInConfig(); err != nil {
			return nil, errors.WrapError(err, errors.Internal, "failed to load configuration file (DEBUG=true)")
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.WrapError(err, errors.Internal, "failed to unmarshal config")
	}

	if err := cfg.Validate(val); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Validate(v *validator.Validate) error {
	v.RegisterStructValidation(validateDatabaseConfig, Config{})

	if err := v.Struct(cfg); err != nil {
		return errors.WrapError(err, errors.Internal, "config validation failed")
	}

	return nil
}

func validateDatabaseConfig(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(Config)

	if cfg.StoreType == "mongo" {
		if cfg.DatabaseURL == "" {
			sl.ReportError(cfg.DatabaseURL, "DatabaseURL", "DatabaseURL", "required_if_store_mongo", "")
		}
		if cfg.DatabaseName == "" {
			sl.ReportError(cfg.DatabaseName, "DatabaseName", "DatabaseName", "required_if_store_mongo", "")
		}
	}
}
