package main

import (
	"os"
	"search-engine/pkg/menu"
	"search-engine/pkg/revindex"
)

func main() {
	reader := os.Stdin
	writer := os.Stdout

	indexes := revindex.InitIndexes()

	for {
		operation := menu.MenuOperation(reader, writer, indexes)
		if operation == 1 {
			revindex.AddIndexesToFile(indexes)
		}
	}
}
