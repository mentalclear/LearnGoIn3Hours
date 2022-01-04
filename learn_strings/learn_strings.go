package main

import "fmt"

func main() {
	// var s string
	// s = "Hello, World!"
	s := "Hello, \n\t\"World!\" with a bakcslash \\" // type inference
	fmt.Println(s)
	// With backticks
	s1 := `Hello, "world! with a backslash \`
	fmt.Println(s1)

	ss1 := "👋 🌎"
	ss2 := "Hello World!"
	ss3 := ss2 + " " + ss1
	fmt.Println(ss3)

	// indexing
	new_s := "Hello, World!"
	new_s2 := new_s[0:5]
	new_s3 := new_s[7:13]
	new_s4 := new_s[:5]
	new_s5 := new_s[7:]

	fmt.Println(new_s, new_s2, new_s3, new_s4, new_s5)
	fmt.Println(len(new_s))

	ss2_1 := ss1[:1]
	ss2_2 := ss1[2:]
	fmt.Println(ss1, len(ss1), ss2_1, len(ss2_1), ss2_2, len(ss2_2))

	// rune
	s3s := "Hello, "
	//var r rune = 127757
	r := '🌎' // rune
	s3s += string(r)
	fmt.Println(s3s)

	// Arrays
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	fmt.Println(vals, vals[0], vals[1], vals[2])

	// vals[3] = 8 generates out of bounds error

	// Cannot assign [4]int to [3]int
	// var vals2 [4]int = vals
	// fmt.Println(vals, vals2)

}
