package leftpad

import (
	"strings"
	"unicode/utf8"
)

var default_char = ' ' // Package variable declarations

// var/func name started with lowercase - package private
// when starting with capital letter - public

// Comments come here
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

// Comments come here
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
