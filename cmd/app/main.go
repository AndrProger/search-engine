package main

import (
	"bufio"
	"fmt"
	"os"
	extractor "search-engine/pkg/extractor/url"
	"search-engine/pkg/revindex"
)

func main() {
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Проиндексировать сайт")
		fmt.Println("2. Поиск по индексам")
		fmt.Println("3. Выход")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ваш выбор: ")
		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\n":
			var url string
			fmt.Print("Введите URL для индексации: ")
			url, _ = reader.ReadString('\n')
			url = url[:len(url)-1] // удаляем символ новой строки
			indexWebsite(url)
		case "2\n":
			var word string
			fmt.Print("Введите слово для поиска: ")
			word, _ = reader.ReadString('\n')
			word = word[:len(word)-1] // удаляем символ новой строки
			searchWord(word)
		case "3\n":
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println("Некорректный выбор. Попробуйте снова.")
		}
	}
}

func indexWebsite(url string) {
	words, err := extractor.ExtractUrl(url)
	if err != nil {
		println("Ошибка при индексации сайта:", err)
	}
	fmt.Println("Сайт", url, "проиндексирован.")
	revindex.AddIndexes(words, url)
}

func searchWord(word string) {
	urls := revindex.GetUrls(word)
	fmt.Printf("Список URL, содержащих слова [%v] : %v", word, urls)
}
