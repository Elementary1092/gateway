package config

import (
	"github.com/elem1092/gateway/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Configuration struct {
	GatewayCfg GatewayConfig        `yaml:"gateway"`
	FetcherCfg FetcherServiceConfig `yaml:"fetcher_service"`
	CRUDCfg    CRUDServiceConfig    `yaml:"crud_service"`
}

type GatewayConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type FetcherServiceConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type CRUDServiceConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

var cfg = &Configuration{GatewayConfig{}, FetcherServiceConfig{}, CRUDServiceConfig{}}
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Parsing configurations")
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Fatalf("Unable to parse configurations due to: %s", help)
			panic(err)
		}
	})

	return cfg
}
