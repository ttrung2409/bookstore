package strings

func Filter(arr []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
