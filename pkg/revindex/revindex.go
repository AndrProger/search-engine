package revindex

var indexes = make(map[string][]string)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func AddIndexes(strArr []string, url string) {
	for _, v := range strArr {
		if !contains(indexes[v], url) {
			indexes[v] = append(indexes[v], url)
		}
	}
}

func GetUrls(str string) []string {
	return indexes[str]
}
