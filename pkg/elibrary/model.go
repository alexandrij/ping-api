package elibrary

import (
	"github.com/Alexandrij/ping-api/pkg/article"
)

type Article = article.Article
type SearchRequest = article.SearchRequest
type SearchResult = article.SearchResult
type Service interface {
	article.Service
}

type Profile struct {
	GUID   string
	UserID string
}
