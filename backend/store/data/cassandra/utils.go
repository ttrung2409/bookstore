package data

import (
	"regexp"
	"strings"
)

func FieldNameToColumnName(field string) string {
	return strings.ToLower(regexp.MustCompile(`(.)([A-Z])`).ReplaceAllString(field, "$1_$2"))
}