package main

import "fmt"

func addNumbers(a int, b int) int {
	return a + b
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFail(a int, arr [2]int, s string) {
	a *= 2
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	s += s
	fmt.Println("in doubleFail: ", a, arr, s)
}

func main() {
	// 	a := addNumbers(2, 3)
	// 	fmt.Println(a)

	// 	b := addNumbers(4, 10)
	// 	fmt.Println(b)

	// 	c := addNumbers(100, -100)
	// 	fmt.Println(c)
	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20)

	a := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Println("in main: ", a, arr, s)

}
