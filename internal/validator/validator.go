package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

// Valid() returns true if the FieldErrors map doesn't contain any entries.
func (v *Validator) Validate() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}
func (v *Validator) AddFieldError(key string, errorMsg string) {
	// essendo un puntatore devo controllare che esita
	if v.FieldErrors == nil {
		v.FieldErrors = map[string]string{}
	}
	// salvo solo la prima occorrenza di errore per chiave, altrimenti si puo fare map[sting][]string
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = errorMsg
	}
}
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt() returns true if a value is in a list of permitted integers.
func PermittedInt(value int, permittedValues ...int) bool {
	for _, item := range permittedValues {
		if item == value {
			return true
		}
	}
	return false
}

// MinChars() returns true if a value contains at least n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

var EmailRX = regexp.MustCompile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)

// Matches() returns true if a value matches a provided compiled regular
// expression pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
