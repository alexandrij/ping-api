package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"ping-api/pkg/logger"
	"strconv"
)

// ELibraryHandler структура для обработчика elibrary
type ELibraryHandler struct {
	// Здесь может быть сервис для работы с elibrary API
}

// NewELibraryHandler создает новый экземпляр обработчика
func NewELibraryHandler() *ELibraryHandler {
	return &ELibraryHandler{}
}

// SearchPublicationsRequest структура запроса для поиска публикаций
type SearchPublicationsRequest struct {
	Query    string `json:"query"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

// SearchPublicationsResponse структура ответа для поиска публикаций
type SearchPublicationsResponse struct {
	Publications []Publication `json:"publications"`
	Total        int           `json:"total"`
	Page         int           `json:"page"`
}

// Publication структура публикации
type Publication struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Authors   []Author `json:"authors"`
	Abstract  string   `json:"abstract"`
	Year      int      `json:"year"`
	DOI       string   `json:"doi"`
	Journal   string   `json:"journal"`
	Citations int      `json:"citations"`
}

// Author структура автора
type Author struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	ORCID string `json:"orcid,omitempty"`
}

// SearchPublications обрабатывает поиск публикаций
func (h *ELibraryHandler) SearchPublications(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	// Установка значений по умолчанию
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	logger.Info("Get publications", zap.String("query", query))

	// Здесь должна быть логика поиска в elibrary
	// Для примера возвращаем заглушку
	response := SearchPublicationsResponse{
		Publications: []Publication{
			{
				ID:    1,
				Title: "Пример научной статьи",
				Authors: []Author{
					{ID: 1, Name: "Иванов И.И."},
					{ID: 2, Name: "Петров П.П.", ORCID: "0000-0000-0000-0000"},
				},
				Abstract:  "Аннотация статьи...",
				Year:      2023,
				DOI:       "10.1234/example",
				Journal:   "Научный журнал",
				Citations: 42,
			},
		},
		Total: 1,
		Page:  page,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetPublicationByID возвращает информацию о публикации по ID
func (h *ELibraryHandler) GetPublicationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid publication ID", http.StatusBadRequest)
		return
	}

	// Здесь должна быть логика получения публикации из elibrary
	// Для примера возвращаем заглушку
	publication := Publication{
		ID:    id,
		Title: "Научная статья с ID " + strconv.Itoa(id),
		Authors: []Author{
			{ID: 1, Name: "Автор 1"},
			{ID: 2, Name: "Автор 2"},
		},
		Abstract:  "Полный текст аннотации...",
		Year:      2023,
		DOI:       "10.1234/example." + strconv.Itoa(id),
		Journal:   "Журнал научных публикаций",
		Citations: 10,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publication)
}

// SearchAuthors обрабатывает поиск авторов
func (h *ELibraryHandler) SearchAuthors(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	logger.Info("Get authors", zap.String("query", query))

	// Здесь должна быть логика поиска авторов в elibrary
	// Для примера возвращаем заглушку
	authors := []Author{
		{ID: 1, Name: "Иванов Иван Иванович", ORCID: "0000-0000-0000-0001"},
		{ID: 2, Name: "Петрова Мария Сергеевна", ORCID: "0000-0000-0000-0002"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// RegisterRoutes регистрирует маршруты для elibrary handler
func (h *ELibraryHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/elibrary/publications/search", h.SearchPublications).Methods("GET")
	router.HandleFunc("/elibrary/publications/{id:[0-9]+}", h.GetPublicationByID).Methods("GET")
	router.HandleFunc("/elibrary/authors/search", h.SearchAuthors).Methods("GET")
}
