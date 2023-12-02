package htmx

import "strings"

func sanitize(s string) string {
	sanitized := strings.ReplaceAll(s, "\n ", "")
	for strings.Contains(sanitized, "\n") {
		sanitized = strings.Replace(sanitized, "\n", "", -1)
	}
	for strings.Contains(sanitized, " <") {
		sanitized = strings.Replace(sanitized, " <", "<", -1)
	}
	for strings.Contains(sanitized, "  ") {
		sanitized = strings.Replace(sanitized, "  ", " ", -1)
	}
	return sanitized
}
