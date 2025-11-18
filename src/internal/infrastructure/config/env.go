package config

import (
	"fmt"
	"golang-web-server/src/internal/infrastructure/appErrors"
	"os"
)

type ConfigService struct
{
	data *Config
}

func NewConfigService() IConfigService {
	return &ConfigService{}
}


func(cs *ConfigService) GetConfig() (*Config, error){
	return cs.loadConfig()
}

func(cs *ConfigService) loadConfig() (*Config, error){
	host, err := cs.loadEnv(APIHost)
	if err != nil{
		return nil, err
	}

	port, err := cs.loadEnv(APIPort)
	if err != nil{
		return nil, err
	}

	env_mode, err := cs.loadEnv(Environment)
	if err != nil{
		return nil, err
	}
	return &Config{
		Server: Server{
			Host: host,
			Port: port,
			Environment: env_mode,
		},
	}, nil
}

func (cs *ConfigService) loadEnv(envName string) (string, error) {
	result := os.Getenv(envName)
	if result == "" {
		return "", appErrors.New(appErrors.EnvError, fmt.Sprintf("%s: %s", appErrors.EnvVarMissingValue, envName))
	}
	return result, nil
}