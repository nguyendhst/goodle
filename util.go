package goodle

import (
	"strings"
)

var (
	MIMETypes = []string{
		"text/html",
		"text/plain",
		"text/xml",
		"text/json",

		"application/xml",
		"application/json",
		"application/xhtml+xml",
		"application/rss+xml",
	}
)

func isNonTextualContentType(contentType string) bool {
	for _, t := range MIMETypes {
		if strings.HasPrefix(contentType, t) {
			return false
		}
	}
	return true
}
