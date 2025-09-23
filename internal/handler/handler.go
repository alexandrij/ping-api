package handler

import (
	"github.com/Alexandrij/ping-api/config"
	"github.com/Alexandrij/ping-api/internal/service"
	"github.com/gorilla/mux"
)

// Handler структура для всех обработчиков
type Handler struct {
	ArticleHandler *ArticleHandler
	// Другие обработчики...
}

// NewHandler создает новый экземпляр Handler
func NewHandler(cfg config.ApiConfig) *Handler {
	// Создаем сервис для работы со статьями
	articleService := service.NewArticleService(cfg.Elibrary.ApiKey, cfg.Arxiv.Endpoint)

	return &Handler{
		ArticleHandler: NewArticleHandler(articleService),
		// Инициализация других обработчиков...
	}
}

// RegisterRoutes регистрирует все маршруты
func (h *Handler) RegisterRoutes(router *mux.Router) {
	h.ArticleHandler.RegisterRoutes(router)
	// Регистрация других маршрутов...
}
