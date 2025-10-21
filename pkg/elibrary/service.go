package elibrary

import (
	"context"
	"fmt"
	"github.com/Alexandrij/ping-api/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ELibraryService интерфейс для работы с elibrary.ru
type ELibraryService interface {
	SearchArticles(ctx context.Context, req SearchRequest) (*SearchResult, error)
	GetArticleByID(ctx context.Context, id string) (*Article, error)
}

// NewELibraryService создает новый сервис для работы с elibrary.ru
func NewELibraryService(eProfile Profile) ELibraryService {
	return &elibraryService{GUID: eProfile.GUID, UserID: eProfile.UserID}
}

type elibraryService struct {
	GUID   string
	UserID string
}

func (e *elibraryService) SearchArticles(ctx context.Context, searchReq SearchRequest) (*SearchResult, error) {
	client := &http.Client{}

	data := url.Values{
		"where_fulltext":    []string{"on"},
		"where_name":        []string{"on"},
		"where_abstract":    []string{"on"},
		"where_keywords":    []string{"on"},
		"where_affiliation": []string{""},
		"where_references":  []string{""},
		"type_article":      []string{"on"},
		"type_disser":       []string{"on"},
		"type_book":         []string{"on"},
		"type_report":       []string{"on"},
		"type_conf":         []string{"on"},
		"type_patent":       []string{"on"},
		"type_preprint":     []string{"on"},
		"type_grant":        []string{"on"},
		"type_dataset":      []string{"on"},
		"search_freetext":   []string{""},
		"search_morph":      []string{"on"},
		"search_fulltext":   []string{""},
		"search_open":       []string{""},
		"search_results":    []string{""},
		"titles_all":        []string{""},
		"authors_all":       []string{""},
		"rubrics_all":       []string{""},
		"queryboxid":        []string{""},
		"itemboxid":         []string{""},
		"begin_year":        []string{""},
		"end_year":          []string{""},
		"issues":            []string{"all"},
		"orderby":           []string{"rank"},
		"order":             []string{"rev"},
		"changed":           []string{"1"},
		"ftext":             []string{"компьютерное зрение"},
	}

	req, err := http.NewRequest("POST", "https://elibrary.ru/query_results.asp", strings.NewReader(data.Encode()))

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 YaBrowser/25.8.0.0 Safari/537.36")

	req.AddCookie(&http.Cookie{Name: "SCookieGUID", Value: e.GUID})
	req.AddCookie(&http.Cookie{Name: "SUserID", Value: e.UserID})

	logger.Debug(req.URL.String(), zap.String("query", data.Encode()), zap.String("cookies", req.Header.Get("Cookie")))

	resp, err := client.Do(req)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error(fmt.Sprintf("elibrary.ru returned status code %d", resp.StatusCode))
		return nil, fmt.Errorf("elibrary.ru returned status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	fmt.Println(string(body), resp.StatusCode, resp.Status)

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
		Page:     searchReq.Page,
		PageSize: searchReq.PageSize,
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
