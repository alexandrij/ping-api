package elibrary

import (
	"net/http"
)

func ApplyCookie(r *http.Request) *http.Request {
	// Добавляем cookies из конфигурации
	cookies := []*http.Cookie{
		{
			Name:  "_ym_uid",
			Value: "1757317899542484987",
		},
		{
			Name:  "_ym_d",
			Value: "1757317899",
		},
		{
			Name:  "__utmz",
			Value: "216042306.1757327149.2.2.utmcsr=elibrary.ru|utmccn=(referral)|utmcmd=referral|utmcct=/",
		},
		{
			Name:  "__utmc",
			Value: "216042306",
		},
		{
			Name:  "_ym_isad",
			Value: "2",
		},
		{
			Name:  "__utma",
			Value: "216042306.1664554439.1757317903.1759327985.1759383774.11",
		},
		{
			Name:  "__utmt",
			Value: "1",
		},
		{
			Name:  "__utmb",
			Value: "216042306.4.10.1759383774",
		},
		{
			Name:  "SCookieGUID",
			Value: "F05511BB-D1D8-4AC6-A4AA-68AF81A83F3A",
		},
		{
			Name:  "SUserID",
			Value: "631204980",
		},
	}

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	return r
}

func ApplyHeaders(r *http.Request) *http.Request {
	// Добавляем headers из конфигурации
	// Устанавливаем заголовки
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Referer", "https://elibrary.ru/querybox.asp?scope=infound")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	r = ApplyCookie(r)

	return r
}
