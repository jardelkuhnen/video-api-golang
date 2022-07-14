package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateBadLanguage(field validator.FieldLevel) bool {
	for i := 0; i < len(GetBadWords()); i++ {
		if strings.Contains(field.Field().String(), GetBadWords()[i]) {
			return false
		}
	}
	return true
}

func GetBadWords() []string {
	return []string{"viado", "gay"}
}
