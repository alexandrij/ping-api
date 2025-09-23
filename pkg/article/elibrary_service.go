package article

import (
	"context"
	"time"
)

// ELibraryService интерфейс для работы с elibrary.ru
type ELibraryService interface {
	SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error)
	GetArticleByID(ctx context.Context, id string) (*Article, error)
}

// NewELibraryService создает новый сервис для работы с elibrary.ru
func NewELibraryService(apiKey string) ELibraryService {
	return &elibraryService{
		apiKey: apiKey,
	}
}

type elibraryService struct {
	apiKey string
}

func (e *elibraryService) SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error) {
	// Здесь должна быть реализация обращения к API elibrary.ru
	// Для примера возвращаем заглушку
	result := &SearchResult{
		Articles: []Article{
			{
				ID:          "elibrary_12345",
				Title:       "Пример статьи из elibrary",
				Authors:     []string{"Иванов И.И.", "Петров П.П."},
				Abstract:    "Аннотация статьи из elibrary...",
				Publication: "Научный журнал",
				Volume:      "Том 15, № 3",
				Date:        time.Now().AddDate(0, -2, 0),
				DOI:         "10.1234/elibrary.12345",
				URL:         "https://elibrary.ru/item.asp?id=12345",
				Source:      "elibrary",
			},
		},
		Total:    1,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return result, nil
}

func (e *elibraryService) GetArticleByID(ctx context.Context, id string) (*Article, error) {
	// Здесь должна быть реализация получения статьи по ID из elibrary.ru
	// Для примера возвращаем заглушку
	article := &Article{
		ID:          id,
		Title:       "Полная информация о статье из elibrary",
		Authors:     []string{"Иванов И.И.", "Петров П.П.", "Сидоров С.С."},
		Abstract:    "Полный текст аннотации статьи из elibrary...",
		Publication: "Научный журнал",
		Volume:      "Том 15, № 3",
		Date:        time.Now().AddDate(0, -2, 0),
		DOI:         "10.1234/elibrary.12345",
		URL:         "https://elibrary.ru/item.asp?id=12345",
		Source:      "elibrary",
	}

	return article, nil
}
