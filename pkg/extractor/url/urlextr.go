package url

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func ExtractUrl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			panic("Error while closing response body: " + cerr.Error())
		}
	}()
	content := extractText(resp.Body)

	return content, nil
}

func extractText(body io.Reader) []string {
	tokenizer := html.NewTokenizer(body)
	var textBuilder strings.Builder

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			// Используем strings.Fields для разделения текста на слова
			return strings.Fields(textBuilder.String())
		case html.TextToken:
			text := strings.TrimSpace(html.UnescapeString(string(tokenizer.Text())))
			if len(text) > 0 {
				textBuilder.WriteString(text + " ") // Добавляем пробел после каждого слова для корректного разделения
			}
		default:
			// На случай, если встретится необработанный тип токена
		}
	}
}
