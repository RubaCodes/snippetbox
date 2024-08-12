package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

// Valid() returns true if the FieldErrors map doesn't contain any entries.
func (v *Validator) Validate() bool {
	return len(v.FieldErrors) == 0
}
func (v *Validator) addFieldError(key string, errorMsg string) {
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
		v.addFieldError(key, message)
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
