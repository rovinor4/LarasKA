package utils

import (
	"regexp"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsNumeric(input string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(input)
}

func RupiahFormat(r int) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("Rp %d", r)
}

func IsIn(value string, options []string) bool {
	for _, option := range options {
		if value == option {
			return true
		}
	}
	return false
}
