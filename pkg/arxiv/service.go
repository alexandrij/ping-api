package arxiv

import (
	"context"
	"time"
)

// ArxivService интерфейс для работы с arxiv.org
type ArxivService interface {
	SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error)
	GetArticleByID(ctx context.Context, id string) (*Article, error)
}

// NewArxivService создает новый сервис для работы с arxiv.org
func NewArxivService(endpoint string) ArxivService {
	if endpoint == "" {
		endpoint = "http://export.arxiv.org/api/query"
	}
	return &arxivService{
		endpoint: endpoint,
	}
}

type arxivService struct {
	endpoint string
}

func (a *arxivService) SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error) {
	// Здесь должна быть реализация обращения к API arxiv.org
	// Для примера возвращаем заглушку
	result := &SearchResult{
		Articles: []Article{
			{
				ID:          "arxiv:2101.12345",
				Title:       "Example article from arXiv",
				Authors:     []string{"Smith, John", "Doe, Jane"},
				Abstract:    "Abstract of the article from arXiv...",
				Publication: "arXiv e-prints",
				Date:        time.Now().AddDate(0, -1, 0),
				DOI:         "10.48550/arXiv.2101.12345",
				URL:         "https://arxiv.org/abs/2101.12345",
				Source:      "arxiv",
			},
		},
		Total:    1,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return result, nil
}

func (a *arxivService) GetArticleByID(ctx context.Context, id string) (*Article, error) {
	// Здесь должна быть реализация получения статьи по ID из arxiv.org
	// Для примера возвращаем заглушку
	article := &Article{
		ID:          id,
		Title:       "Full information about article from arXiv",
		Authors:     []string{"Smith, John", "Doe, Jane", "Brown, Alice"},
		Abstract:    "Full abstract text of the article from arXiv...",
		Publication: "arXiv e-prints",
		Date:        time.Now().AddDate(0, -1, 0),
		DOI:         "10.48550/arXiv.2101.12345",
		URL:         "https://arxiv.org/abs/2101.12345",
		Source:      "arxiv",
	}

	return article, nil
}
