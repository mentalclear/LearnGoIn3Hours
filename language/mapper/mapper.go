package mapper

import (
	"github.com/rylans/getlang"
)

func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "de":
		return "Guten Tag!"
	case "fr":
		return "Bonjour!"
	default:
		return "IDK your language yet!"
	}
}
