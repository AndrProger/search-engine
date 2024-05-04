package url

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func ExtractUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil {
			err = cerr // Присваиваем ошибку внешней переменной err
		}
	}(resp.Body)
	content := extractText(resp.Body)
	if err != nil {
		return "", err // Возвращаем ошибку извлечения текста
	}
	return content, nil
}

func extractText(body io.Reader) string {
	tokenizer := html.NewTokenizer(body)
	var textBuilder strings.Builder

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return textBuilder.String()
		case html.TextToken:
			text := strings.TrimSpace(html.UnescapeString(string(tokenizer.Text())))
			if len(text) > 0 {
				textBuilder.WriteString(text)
				textBuilder.WriteString("\n")
			}
		}
	}
}
