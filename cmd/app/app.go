package app

import (
	"fmt"
	"github.com/Alexandrij/ping-api/config"
	"github.com/Alexandrij/ping-api/internal/handler"
	"github.com/Alexandrij/ping-api/pkg/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
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

	// Создаем обработчики
	h := handler.NewHandler(cfg.Api)

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
