package main

import (
	"fmt"

	"github.com/mentalclear/LearnGoIn3Hours/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("Hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("¿Cómo está?", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
	fmt.Println(leftpad.FormatRune("¿Cómo está?", 15, '😀'))
}
