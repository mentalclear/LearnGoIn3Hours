package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(float32(i) + f)
	fmt.Println(i + int8(f))     // Truncates f to an int8
	fmt.Println(i + int8(f+1.8)) // math in type conversions

	// more mismatches
	var i2 int8 = 20
	var j2 int32 = 40
	fmt.Println(int32(i2) + j2)

}
