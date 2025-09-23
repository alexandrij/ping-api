package config

import (
	"fmt"
	"github.com/Alexandrij/ping-api/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Server ServerConfig
		Api    ApiConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	ElibraryConfig struct {
		ApiKey string
	}

	ArxivConfig struct {
		Endpoint string
	}

	ApiConfig struct {
		Elibrary ElibraryConfig
		Arxiv    ArxivConfig
	}
)

func Load() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("./configs")

	// Автоматическая привязка переменных окружения
	viper.AutomaticEnv()

	// Установка приоритета для локальных значений из .env
	viper.SetConfigName("app")
	// Объединяем с основной конфигурацией

	if err := viper.MergeInConfig(); err != nil {
		logger.Error("Error reading config file", zap.Error(err))
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file", zap.Error(err))
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{
		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		Api: ApiConfig{
			Elibrary: ElibraryConfig{
				ApiKey: viper.GetString("api.elibrary.apiKey"),
			},
			Arxiv: ArxivConfig{
				Endpoint: viper.GetString("api.arxiv.endpoint"),
			},
		},
	}

	logger.Info("Config loaded", zap.Any("config", config))

	return config, nil
}
