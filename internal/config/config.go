package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const DefaultConfigPath = "../configs/server-config.yml"

type HTTPServerConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	LogLevel string `yaml:"log_level" env-default:"debug"`
}

func MustLoad(configPath string) *HTTPServerConfig {
	//TODO добавить функциональность чтения из параметров
	//TODO добавить проверку на октрытие файла

	var ServerConfig HTTPServerConfig

	if err := cleanenv.ReadConfig(configPath, &ServerConfig); err != nil {
		log.Fatalf("Ошибка при прочтении файла конфигурации: %s", err)
	}

	return &ServerConfig
}

func ParseAdress(config *HTTPServerConfig) string {
	return config.Host + ":" + config.Port
}
