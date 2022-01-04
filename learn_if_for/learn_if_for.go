package main

import "fmt"

func ifs() {
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}
}

func for_looping() {
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	b := 3
	for j := 0; j < 10; j++ {
		if j == b {
			continue
		}
		fmt.Println(j)
	}
}

func like_while() {
	// While analog
	i2 := 0
	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}
	fmt.Println(i2)
}

func like_while_true() {
	i := 0
	for {
		fmt.Println(i)
		i += 1
		if i > 10 {
			break
		}
	}
	fmt.Println(i)
}

func for_range_loop(s string) {
	//s := "Hello, World!"
	for k, v := range s { // Here v is the rune
		fmt.Println(k, v, string(v))
	}
}

func main() {

	// ifs()
	//for_looping()
	//like_while()
	// like_while_true()

	for_range_loop("Hello, World!")
	fmt.Println()
	for_range_loop("👋 🌎")
}
