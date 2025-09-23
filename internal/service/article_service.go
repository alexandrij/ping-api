package service

import (
	"context"
	"github.com/Alexandrij/ping-api/pkg/article"
	"github.com/Alexandrij/ping-api/pkg/arxiv"
	"github.com/Alexandrij/ping-api/pkg/elibrary"
)

type Article = article.Article

// SearchRequest запрос для поиска статей
type SearchRequest = article.SearchRequest

// SearchResult результат поиска статей
type SearchResult = article.SearchResult

// Service интерфейс для работы со статьями
type Service interface {
	SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error)
	GetArticleByID(ctx context.Context, id string, source string) (*Article, error)
}

// NewService создает новый сервис для работы со статьями
func NewArticleService(elibraryKey, arxivEndpoint string) Service {
	return &service{
		elibrary: elibrary.NewELibraryService(elibraryKey),
		arxiv:    arxiv.NewArxivService(arxivEndpoint),
	}
}

type service struct {
	elibrary elibrary.ELibraryService
	arxiv    arxiv.ArxivService
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
