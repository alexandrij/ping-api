package article

import (
	"context"
	"time"
)

// Article представляет научную статью
type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Authors     []string  `json:"authors"`
	Abstract    string    `json:"abstract"`
	Publication string    `json:"publication"`
	Volume      string    `json:"volume,omitempty"`
	Date        time.Time `json:"date"`
	DOI         string    `json:"doi,omitempty"`
	URL         string    `json:"url"`
	Source      string    `json:"source"` // "elibrary" или "arxiv"
}

// SearchRequest запрос для поиска статей
type SearchRequest struct {
	Query    string `json:"query"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Source   string `json:"source"` // "elibrary", "arxiv" или "all"
}

// SearchResult результат поиска статей
type SearchResult struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

// Service интерфейс для работы со статьями
type Service interface {
	SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error)
	GetArticleByID(ctx context.Context, id string, source string) (*Article, error)
}
