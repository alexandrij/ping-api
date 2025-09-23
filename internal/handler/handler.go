package handler

import (
	"ping-api/pkg/article"

	"github.com/gorilla/mux"
)

// Handler структура для всех обработчиков
type Handler struct {
	ArticleHandler *ArticleHandler
	// Другие обработчики...
}

// NewHandler создает новый экземпляр Handler
func NewHandler(elibraryKey, arxivEndpoint string) *Handler {
	// Создаем сервис для работы со статьями
	articleService := article.NewService(elibraryKey, arxivEndpoint)

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
