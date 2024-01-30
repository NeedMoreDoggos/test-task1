package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_PATH_ENV     = "CONFIG_PATH"
	DEFAULT_CONFIG_PATH = "./config.yaml"
)

var ENV_TYPES = map[string]bool{
	"local": true,
	"dev":   true,
	"prod":  true,
}

var (
	ErrParseYaml          = errors.New("failed to parse yaml")
	ErrParseJSON          = errors.New("failed to parse json")
	ErrUndefinedExtension = errors.New("undefined extension")
	ErrReadFile           = errors.New("failed to read file")
	ErrValidateConfig     = errors.New("failed to validate config")
)

func MustConfig() *Config {
	filePath := os.Getenv(CONFIG_PATH_ENV)
	if filePath == "" {
		filePath = DEFAULT_CONFIG_PATH
	}

	log.Printf("try to read config from %s", filePath)

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("config file not found: %s", err.Error()))
		}
	}

	config, err := parseConfig(filePath)
	if err != nil {
		panic(err)
	}

	log.Printf("config successfully read from %s", filePath)

	err = validateConfig(config)
	if err != nil {
		panic(err)
	}

	return config
}

func parseConfig(filePath string) (*Config, error) {
	extension := filepath.Ext(filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, ErrReadFile
	}

	log.Printf("read config from %s", filePath)
	log.Printf("extension: %s", extension)

	switch extension {
	case ".yaml", ".yml":
		return parseYaml(file)
	case ".json":
		return parseJson(file)
	default:
		return nil, ErrUndefinedExtension
	}
}

func parseYaml(file []byte) (*Config, error) {
	var cfg *Config

	err := yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, ErrParseYaml
	}

	return cfg, nil
}

func parseJson(file []byte) (*Config, error) {
	var cfg *Config

	err := json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, ErrParseJSON
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if err := validateDBConfig(cfg.DBConfig); err != nil {
		return err
	}

	if err := validateEnviroment(cfg.Environment); err != nil {
		return err
	}

	return nil
}

func validateDBConfig(cfg DBConfig) error {
	if cfg.Type == "" {
		return fmt.Errorf("%w: type is required", ErrValidateConfig)
	}

	if cfg.Host == "" {
		return fmt.Errorf("%w: host is required", ErrValidateConfig)
	}

	if cfg.Port == "" {
		return fmt.Errorf("%w: port is required", ErrValidateConfig)
	}

	if portNum, err := strconv.ParseUint(cfg.Port, 10, 16); err != nil || portNum == 0 {
		return fmt.Errorf("%w: invalid port: %s", ErrValidateConfig, cfg.Port)
	}

	if cfg.User == "" {
		return fmt.Errorf("%w: user is required", ErrValidateConfig)
	}

	if cfg.Database == "" {
		return fmt.Errorf("%w: database is required", ErrValidateConfig)
	}

	return nil
}

func validateEnviroment(env string) error {
	if !ENV_TYPES[env] {
		return fmt.Errorf("%w: invalid env %s", ErrValidateConfig, env)
	}

	return nil
}
