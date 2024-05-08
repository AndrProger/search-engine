package revindex

import (
	"encoding/json"
	"io"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func AddIndexes(indexes map[string][]string, strArr []string, url string) map[string][]string {
	for _, v := range strArr {
		if !contains(indexes[v], url) {
			indexes[v] = append(indexes[v], url)
		}
	}
	return indexes
}

func WriteIndexesToJson(indexes map[string][]string, w io.Writer) error {
	jsonData, err := json.Marshal(indexes)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func ReadIndexesFromJson(r io.Reader) (map[string][]string, error) {
	var indexes map[string][]string
	err := json.NewDecoder(r).Decode(&indexes)
	if err != nil {
		return nil, err
	}
	return indexes, nil
}
