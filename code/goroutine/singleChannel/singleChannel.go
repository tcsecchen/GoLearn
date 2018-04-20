package main

import (
	"fmt"
)

func counter(out chan<- int) {
	defer close(out)
	for x := 0; x < 100; x++ {
		out <- x
	}
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for v := range in {
		out <- v
	}
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
