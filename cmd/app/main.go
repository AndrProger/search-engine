package main

import (
	"fmt"
	"net/http"
)

func main() {
	// URL сайта, который вы хотите получить
	url := "https://habr.com/ru/articles/812199/"

	// Выполнение GET-запроса к сайту
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Извлечение текстового содержимого из HTML
	text := ExtractText(resp.Body)

	// Вывод текста
	fmt.Println(text)
}

// ExtractText извлекает текст из HTML
