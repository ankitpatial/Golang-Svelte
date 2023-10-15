package slice

import "strings"

func Unique[T any](arr []T) []T {
	keys := make(map[any]bool)
	var list []T
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// HasFold checks if the given string(case-insensitivity) is in the slice
func HasFold(s string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(s, v) {
			return true
		}
	}

	return false
}
