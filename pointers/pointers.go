package main

import "fmt"

func main() {
	a := 10
	b := &a // b is now pinter to a
	c := a
	// b here prents the address, then *b prents value(dereference the pointer)
	fmt.Println(a, b, *b, c) // 10 ADDRESS 10 10

	a = 20
	fmt.Println(a, b, *b, c) // 20 SAME!ADDRESS 20 10 (c is a copy, not a reference)

	*b = 30
	fmt.Println(a, b, *b, c) // 30 SAME!ADDRESS 30 10 value changed on pointer

	c = 40
	fmt.Println(a, b, *b, c) // 30 SAME1ADDRESS 30 40 previous changes persist, only c changed

	// Pointer when not dereferenced returns the address in memory!
	fmt.Println()

	// var b1 *int
	b1 := new(int) // Another way to create a pointer
	fmt.Println(b1)
	fmt.Println(*b1)

	fmt.Println("-----------")
	// Pointers an functions:
	a1 := 20
	fmt.Println(a1)
	setTo10(&a1)
	fmt.Println(a1)

	fmt.Println("-----------")
	a2 := 20
	fmt.Println(a2)
	setTo10Fail(&a2)
	fmt.Println(a2)

}

func setTo10(a *int) {
	*a = 10
}

func setTo10Fail(a *int) {
	a = new(int)
	*a = 10
}
