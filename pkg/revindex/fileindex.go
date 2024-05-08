package revindex

import "os"

func InitIndexes() map[string][]string {
	file, err := os.Open("./indexes.json")
	if err != nil {
		println("Ошибка при открытие файла индексов:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Ошибка при закрытии файла индексов:", err)
		}
	}(file)

	indexes, err := ReadIndexesFromJson(file)
	if err != nil {
		println("Ошибка при создание индекса из файла индексов:", err)
		return make(map[string][]string)
	}
	return indexes
}
func AddIndexesToFile(indexes map[string][]string) {
	file, err := os.Create("./indexes.json")
	if err != nil {
		println("Ошибка при открытие файла индексов:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("Ошибка при закрытии файла индексов:", err)
		}
	}(file)

	err = WriteIndexesToJson(indexes, file)
	if err != nil {
		println("Ошибка при записи индексов в файл:", err)
	}
}
