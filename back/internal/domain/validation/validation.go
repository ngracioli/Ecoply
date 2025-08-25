package validation

import (
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("containsuppercase", containsUppercase)
		v.RegisterValidation("containslowercase", containsLowercase)
		v.RegisterValidation("containsdigit", containsDigit)
		v.RegisterValidation("containsspecial", containsSpecial)
	}
}

func containsUppercase(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, char := range value {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowercase(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, char := range value {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsDigit(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, char := range value {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func containsSpecial(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	specialChars := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~` + "`" + `]`)
	return specialChars.MatchString(value)
}
