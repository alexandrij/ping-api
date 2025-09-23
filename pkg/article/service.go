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

// NewService создает новый сервис для работы со статьями
func NewService(elibraryKey, arxivEndpoint string) Service {
	return &service{
		elibrary: NewELibraryService(elibraryKey),
		arxiv:    NewArxivService(arxivEndpoint),
	}
}

type service struct {
	elibrary ELibraryService
	arxiv    ArxivService
}

func (s *service) SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error) {
	// Установка значений по умолчанию
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.Source == "" {
		req.Source = "all"
	}

	switch req.Source {
	case "elibrary":
		return s.elibrary.SearchArticles(ctx, req)
	case "arxiv":
		return s.arxiv.SearchArticles(ctx, req)
	default: // all
		// Поиск в обоих источниках
		elibReq := req
		elibReq.Source = "elibrary"
		arxivReq := req
		arxivReq.Source = "arxiv"

		elibResult, err := s.elibrary.SearchArticles(ctx, elibReq)
		if err != nil {
			return nil, err
		}

		arxivResult, err := s.arxiv.SearchArticles(ctx, arxivReq)
		if err != nil {
			return nil, err
		}

		// Объединение результатов
		result := &SearchResult{
			Articles: append(elibResult.Articles, arxivResult.Articles...),
			Total:    elibResult.Total + arxivResult.Total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}

		return result, nil
	}
}

func (s *service) GetArticleByID(ctx context.Context, id string, source string) (*Article, error) {
	switch source {
	case "elibrary":
		return s.elibrary.GetArticleByID(ctx, id)
	case "arxiv":
		return s.arxiv.GetArticleByID(ctx, id)
	default:
		// Попробуем найти в обоих источниках
		if article, err := s.elibrary.GetArticleByID(ctx, id); err == nil {
			return article, nil
		}
		return s.arxiv.GetArticleByID(ctx, id)
	}
}
