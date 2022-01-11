package main

import "fmt"

func main() {
	in := make(chan int) // possible solution (chan int, 2) not recommended
	out := make(chan int)

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	in <- 1
	o1 := <-out // This way it won't block
	in <- 2
	//o1 := <-out // A deadlock state when it's here
	o2 := <-out
	fmt.Println(o1, o2)
}
