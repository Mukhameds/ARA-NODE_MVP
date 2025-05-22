
package core

import (
	"strings"
	// другие импорты...
)



// RemoveTag удаляет указанный тег из слайса тегов
func RemoveTag(tags []string, target string) []string {
	var result []string
	for _, tag := range tags {
		if tag != target {
			result = append(result, tag)
		}
	}
	return result
}

// PhaseClose возвращает true, если фазы близки с учётом допуска
func PhaseClose(p1, p2, tolerance float64) bool {
	diff := p1 - p2
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}


// Contains проверяет, содержит ли срез строку
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}


// AddUniqueTag добавляет тег, если его ещё нет в срезе
func AddUniqueTag(tags []string, newTag string) []string {
	for _, t := range tags {
		if t == newTag {
			return tags
		}
	}
	return append(tags, newTag)
}


// Tokenize разбивает строку на слова и приводит их к нижнему регистру
func Tokenize(s string) []string {
	return strings.Fields(strings.ToLower(s))
}
