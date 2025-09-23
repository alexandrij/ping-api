package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"ping-api/config"
	"ping-api/internal/handler"
	"ping-api/pkg/logger"
)

func init() {
	// Инициализация логирования
	logger.Init()
	defer logger.Sync()
}

func Run() error {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Получаем конфигурацию из переменных окружения
	elibraryKey := os.Getenv("ELIBRARY_API_KEY")
	arxivEndpoint := os.Getenv("ARXIV_API_ENDPOINT")

	if arxivEndpoint == "" {
		arxivEndpoint = "http://export.arxiv.org/api/query"
	}

	// Создаем обработчики
	h := handler.NewHandler(cfg.Api.Elibrary.ApiKey, arxivEndpoint)

	// Создаем маршрутизатор
	router := mux.NewRouter()

	// Регистрируем маршруты
	h.RegisterRoutes(router)

	// Запуск сервера
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Starting server",
		zap.String("host", cfg.Server.Host),
		zap.String("port", cfg.Server.Port))

	return http.ListenAndServe(addr, router)
}
