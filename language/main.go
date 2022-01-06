package main

import (
	"fmt"

	"github.com/mentalclear/LearnGoIn3Hours/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}
