package main

import "fmt"

func main() {
	//in := make(chan int) // initial state
	in := make(chan int, 1) // buffered
	out := make(chan int, 1)

	out <- 1

	// in <- 2
	// fmt.Println("wrote 2 to in")
	// v := <-out
	// fmt.Println("read", v, "from out")
	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}

}
