package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

func anotherInterface() {
	var i interface{} = "Hello"
	// i = "Hello"   to hide the warning moving this up to interface declaration
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(j, k, ok)

	// When there is no check, somehting like this will get program to
	// panic state. Look for the ok string above.
	// 	m := i.(int)
	// 	fmt.Println(m)

}

func doStuff(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("Double i is", i+i)
	case string:
		fmt.Println("i is", len(i), "characters long")
	case bool:
		fmt.Println("Tell me the \"true\" words, don't be a \"false\" one")
	default:
		fmt.Println("I don't know what to do with this")
	}
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5)})
	fmt.Println("Result:", result)
	fmt.Println("-----------------")
	anotherInterface()
	fmt.Println("-----------------")
	doStuff(10)
	doStuff("ditch")
	doStuff(true)

}
