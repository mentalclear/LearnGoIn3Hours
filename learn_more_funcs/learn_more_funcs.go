package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a * b
	}
}

func main() {
	// myAddOne := addOne
	// fmt.Println(addOne(1))

	// This is how we declare a func inside another func
	myAddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne(1))
	fmt.Println()
	printOperation(1, addOne)
	printOperation(1, addTwo)

	b := 2
	myAddB := func(a int) int {
		return a + b
	}
	fmt.Println(myAddB(1))

	changeMyB := func(a int) {
		b = a + b
	}
	changeMyB(1)
	fmt.Println(b)
	fmt.Println("----------")

	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(3))
	fmt.Println(addTwo(3))

}
