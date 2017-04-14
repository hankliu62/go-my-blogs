package utils

import (
	"strings"
)

// NormalizeOrderBy trim space for order fields
func NormalizeOrderBy(orderbys []string) []string {
	var sortFields []string
	for _, sortField := range orderbys {
		sortFields = append(sortFields, strings.TrimSpace(sortField))
	}
	return sortFields
}
