package cassandra

import (
	"regexp"
	"strings"
)

func ConvertToColumnNames(fields []string) []string {
	columns := make([]string, len(fields))

	for _, field := range fields {
		columns = append(
			columns,
			strings.ToLower(regexp.MustCompile(`(.)([A-Z])`).ReplaceAllString(field, "$1_$2")),
		)
	}

	return columns
}
