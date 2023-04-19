package utils

import (
	"strings"
	"unicode"

	"github.com/labstack/gommon/bytes"
	"github.com/pkg/errors"
)

// Trim string between two substrings and return the string without it and substrings.
func TrimStringBetween(str, start, end string) string {
	indx1 := strings.Index(str, start)
	indx2 := strings.Index(str, end)
	if indx1 == -1 || indx2 == -1 {
		return strings.TrimSpace(str)
	}
	return strings.TrimSpace(str[:indx1] + str[indx2+len(end):])
}

// TrimRightZeros trims trailing zeros from string.
func TrimRightZeros(str string) string {
	return strings.TrimRight(str, "0")
}

// Parse file size from string to int64
func ParseFileSize(s string) (int64, error) {
	size, err := bytes.Parse(s)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse file size")
	}

	return size, nil
}

// UcFirst capitalizes first letter of a string
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

// Capitalize capitalizes all words in a string
func Capitalize(s string) string {
	var words []string
	for _, word := range strings.Split(s, " ") {
		words = append(words, UcFirst(word))
	}

	return strings.Join(words, " ")
}

// ToSnakeCase converts string to snake case.
func ToSnakeCase(s string) string {
	var snakeCase string
	var lastChar rune
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			if lastChar != '_' {
				snakeCase += "_"
				lastChar = '_'
			}
			continue
		}
		if unicode.IsUpper(r) {
			if i > 0 && lastChar != '_' && !unicode.IsUpper(lastChar) {
				snakeCase += "_"
			}
			snakeCase += string(unicode.ToLower(r))
		} else {
			snakeCase += string(r)
		}
		lastChar = r
	}
	return strings.Trim(snakeCase, "_")
}

// ToTitleCase converts string to title case.
func ToTitleCase(s string) string {
	return Capitalize(strings.ReplaceAll(s, "_", " "))
}
