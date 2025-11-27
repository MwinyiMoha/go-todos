package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/mwinyimoha/commons/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	AppName      string `mapstructure:"APP_NAME" validate:"required"`
	AppVersion   string `mapstructure:"APP_VERSION" validate:"required"`
	AppTimeout   int    `mapstructure:"APP_TIMEOUT" validate:"required,gt=0"`
	Debug        bool   `mapstructure:"DEBUG"`
	ServerPort   int    `mapstructure:"SERVER_PORT"`
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	Store        string `mapstructure:"STORE" validate:"required,oneof=inmemory database"`
	Interface    string `mapstructure:"INTERFACE" validate:"required,oneof=http cli"`
}

func New(val *validator.Validate) (*Config, error) {
	v := viper.New()
	v.SetConfigType("env")

	v.SetDefault("APP_NAME", "")
	v.SetDefault("APP_VERSION", "0.1.0")
	v.SetDefault("DEBUG", true)
	v.SetDefault("APP_TIMEOUT", 10)
	v.SetDefault("SERVER_PORT", 8080)
	v.SetDefault("DATABASE_URL", "")
	v.SetDefault("DATABASE_NAME", "")
	v.SetDefault("STORE", "inmemory")
	v.SetDefault("INTERFACE", "cli")

	v.AutomaticEnv()

	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, errors.WrapError(err, errors.Internal, "failed to load configuration file")
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.WrapError(err, errors.Internal, "failed to unmarshal config")
	}

	val.RegisterStructValidation(validateDependentFields, Config{})
	if err := val.Struct(cfg); err != nil {
		return nil, errors.WrapError(err, errors.InvalidArgument, "config validation failed")
	}

	return &cfg, nil
}

func validateDependentFields(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(Config)

	if cfg.Store == "database" {
		if cfg.DatabaseURL == "" {
			sl.ReportError(cfg.DatabaseURL, "DatabaseURL", "DatabaseURL", "required_if_database_store", "")
		}
		if cfg.DatabaseName == "" {
			sl.ReportError(cfg.DatabaseName, "DatabaseName", "DatabaseName", "required_if_database_store", "")
		}
	}

	if cfg.Interface == "http" {
		if cfg.ServerPort <= 0 || cfg.ServerPort > 65535 {
			sl.ReportError(cfg.ServerPort, "ServerPort", "ServerPort", "required_if_http_interface", "")
		}
	}
}
