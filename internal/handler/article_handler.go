package handler

import (
	"encoding/json"
	"github.com/Alexandrij/ping-api/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// ArticleHandler структура для обработчика статей
type ArticleHandler struct {
	service service.Service
}

// NewArticleHandler создает новый экземпляр обработчика статей
func NewArticleHandler(service service.Service) *ArticleHandler {
	return &ArticleHandler{
		service: service,
	}
}

// SearchArticles обрабатывает поиск статей
func (h *ArticleHandler) SearchArticles(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	source := r.URL.Query().Get("source")

	req := service.SearchRequest{
		Query:    query,
		Page:     page,
		PageSize: pageSize,
		Source:   source,
	}

	response, err := h.service.SearchArticles(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to search articles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetArticleByID возвращает информацию о статье по ID
func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	source := r.URL.Query().Get("source")

	article, err := h.service.GetArticleByID(r.Context(), id, source)
	if err != nil {
		http.Error(w, "Failed to get article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if article == nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// RegisterRoutes регистрирует маршруты для обработчика статей
func (h *ArticleHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/articles/search", h.SearchArticles).Methods("GET")
	router.HandleFunc("/api/articles/{id}", h.GetArticleByID).Methods("GET")
}
