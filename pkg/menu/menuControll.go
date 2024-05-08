package menu

import (
	"bufio"
	"io"
	"os"
	extractor "search-engine/pkg/extractor/url"
	"search-engine/pkg/revindex"
	"strings"
)

func writeMenu(w io.Writer) error {
	s := textMenu()

	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func writeMessage(w io.Writer, s string) {
	_, err := w.Write([]byte(s))
	if err != nil {
		println("Ошибка при выводе сообщения:", err)
		return
	}
}

func textMenu() string {
	return "\nВыберите действие:" +
		"\n1. Проиндексировать сайт" +
		"\n2. Поиск по индексам" +
		"\n3. Выход" +
		"\n--------------------" +
		"\n Ваш выбор:"
}

func readOperation(reader io.Reader) (string, error) {
	readerBuf := bufio.NewReader(reader)

	readString, err := readerBuf.ReadString('\n')
	if err != nil {
		return "", err
	}
	return readString, nil
}

func siteIndexingOperation(reader io.Reader, writer io.Writer, indexes map[string][]string) {
	writeMessage(writer, "Введите URL для индексации: ")

	url, err := readOperation(reader)
	if err != nil {
		println("Ошибка при чтении URL:", err)
		return
	}
	url = url[:len(url)-1]

	words, err := extractor.ExtractUrl(url)
	if err != nil {
		println("Ошибка при индексации сайта:", err)
	}
	writeMessage(writer, "Сайт "+url+" проиндексирован.")
	revindex.AddIndexes(indexes, words, url)
}
func searchWordOperation(reader io.Reader, writer io.Writer, indexes map[string][]string) {
	writeMessage(writer, "Введите слово для поиска: ")

	word, err := readOperation(reader)
	if err != nil {
		println("Ошибка при чтении URL:", err)
		return
	}
	word = word[:len(word)-1]

	if indexes[word] == nil {
		writeMessage(writer, "Слово ["+word+"] не найдено.")
		return
	} else {
		printText := "Список URL, содержащих слова [" + word + "] : " + strings.Join(indexes[word], ", ")
		writeMessage(writer, printText)
	}

}
func defaultOperation(writer io.Writer) {
	writeMessage(writer, "Некорректный выбор. Попробуйте снова.")
}

func MenuOperation(reader io.Reader, writer io.Writer, indexes map[string][]string) int {
	err := writeMenu(os.Stdout)
	if err != nil {
		println("Ошибка при выводе меню:", err)
		return 0
	}

	operation, err := readOperation(os.Stdin)
	if err != nil {
		println("Ошибка при чтении операции:", err)
		return 0
	}
	switch operation {
	case "1\n":
		siteIndexingOperation(reader, writer, indexes)
		return 1
	case "2\n":
		searchWordOperation(reader, writer, indexes)
		return 2
	default:
		defaultOperation(writer)
		return 0
	}
}
