package main

import (
	"fmt"
	"time"
)

func runMe() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go runMe()
	time.Sleep(time.Second)
}
