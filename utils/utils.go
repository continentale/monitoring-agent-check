package utils

func ArrayContains(search string, arr []string) bool {
	for _, value := range arr {
		if search == value {
			return true
		}
	}
	return false
}
