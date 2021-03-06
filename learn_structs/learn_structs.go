package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	Foo
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)
	fmt.Println(h.b)

	fmt.Println("---------------------")
	structTwo()

	fmt.Println("---------------------")
	structThree()

	fmt.Println("---------------------")
	structFour()
}

func structFour() {
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}

func structTwo() {
	f := Foo{
		A: 20,
	}
	var f2 Foo = f
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(f.A)

	var f3 *Foo = &f
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(f.A)
}

func structThree() {
	f := Foo{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F.A)
	b2 := Baz{D: "Nancy", Foo: f}
	fmt.Println(b2.A)

	var f2 Foo = b2.Foo
	fmt.Println(f2)
}
