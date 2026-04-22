package service

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertingData(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("the data cannot be empty")
	}
	isText := strings.ContainsFunc(s, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	})

	if isText {
		return morse.ToMorse(s), nil
	}
	return morse.ToText(s), nil
}
