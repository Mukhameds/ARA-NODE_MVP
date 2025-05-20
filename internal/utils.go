package internal

import "strings"

// ContainsAny проверяет, содержит ли хотя бы один элемент массива одну из подстрок
func ContainsAny(list []string, keywords []string) bool {
	for _, item := range list {
		for _, kw := range keywords {
			if kw != "" && strings.Contains(item, kw) {
				return true
			}
		}
	}
	return false
}
