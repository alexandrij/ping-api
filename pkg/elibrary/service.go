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
	// Здесь должна быть реализация обращения к API elibrary.ru
	// Для примера возвращаем заглушку

	//fetch("https://www.elibrary.ru/query_results.asp",
	//{
	//	"headers": {
	//	"content-type": "application/x-www-form-urlencoded",
	//		"sec-ch-ua": "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"YaBrowser\";v=\"25.8\", \"Yowser\";v=\"2.5\"",
	//		"sec-ch-ua-mobile": "?0",
	//		"sec-ch-ua-platform": "\"Linux\"",
	//		"upgrade-insecure-requests": "1"
	//},
	//	"referrer": "https://www.elibrary.ru/querybox.asp?scope=infound",
	//	"body": "querybox_name=&authors_all=&titles_all=&rubrics_all=&changed=0&queryid=&ftext=%D0%A1%D0%B5%D0%BD%D0%BD%D0%B0%D1%8F+%D0%BF%D0%B0%D0%BB%D0%BE%D1%87%D0%BA%D0%B0&where_name=on&where_abstract=on&where_fulltext=on&where_keywords=on&where_references=&type_article=on&type_disser=on&type_book=on&type_report=on&type_conf=on&type_patent=on&type_preprint=on&type_grant=on&type_dataset=on&search_itemboxid=&search_morph=on&search_results=on&begin_year=0&end_year=0&issues=all&orderby=rank&order=rev&queryboxid=0&save_queryboxid=0",
	//	"method": "POST",
	//	"mode": "cors",
	//	"credentials": "omit"
	//})
	//
	//fetch("https://www.elibrary.ru/query_results.asp",
	//{
	//	"headers": {
	//	"content-type": "application/x-www-form-urlencoded",
	//		"sec-ch-ua": "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"YaBrowser\";v=\"25.8\", \"Yowser\";v=\"2.5\"",
	//		"sec-ch-ua-mobile": "?0",
	//		"sec-ch-ua-platform": "\"Linux\"",
	//		"upgrade-insecure-requests": "1",
	//		"Referer": "https://www.elibrary.ru/querybox.asp?scope=infound"
	//},
	//	"body": "querybox_name=&authors_all=&titles_all=&rubrics_all=&changed=0&queryid=&ftext=%D0%A1%D0%B5%D0%BD%D0%BD%D0%B0%D1%8F+%D0%BF%D0%B0%D0%BB%D0%BE%D1%87%D0%BA%D0%B0&where_name=on&where_abstract=on&where_fulltext=on&where_keywords=on&where_references=&type_article=on&type_disser=on&type_book=on&type_report=on&type_conf=on&type_patent=on&type_preprint=on&type_grant=on&type_dataset=on&search_itemboxid=&search_morph=on&search_results=on&begin_year=0&end_year=0&issues=all&orderby=rank&order=rev&queryboxid=0&save_queryboxid=0",
	//	"method": "POST"
	//})

	//_ym_uid = 1757317899542484987
	//_ym_d = 1757317899
	//__utmz = 216042306.1757327149
	//.2
	//.2.utmcsr = elibrary.ru | utmccn=(referral) | utmcmd = referral | utmcct=/
	//__utmc = 216042306
	//SCookieGUID = F05511BB % 2
	//DD1D8 % 2
	//D4AC6 % 2
	//DA4AA % 2
	//D68AF81A83F3A
	//SUserID = 631204980
	//_ym_isad = 2
	//__utma = 216042306.1664554439
	//.1757317903
	//.1759327985
	//.1759383774
	//.11
	//__utmt = 1
	//__utmb = 216042306.4
	//.10
	//.1759383774

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
		"ftext":             []string{"концепция+общественной+безопасности"},
	}

	req, err := http.NewRequest("POST", "https://www.elibrary.ru/query_results.asp", strings.NewReader(data.Encode()))

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://www.elibrary.ru/querybox.asp?scope=infound")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	req.AddCookie(&http.Cookie{Name: "SCookieGUID", Value: e.GUID})
	req.AddCookie(&http.Cookie{Name: "SUserID", Value: e.UserID})

	logger.Debug(req.URL.String(), zap.String("query", data.Encode()))

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
