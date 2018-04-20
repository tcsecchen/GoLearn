package main

import (
	"fmt"
)

type namer interface {
	area() int
}

type rect struct {
	width, height int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.height * r.width
}

func (s square) area() int {
	return s.side * s.side
}

func main() {
	var a = rect{4, 3}
	var b = square{6}

	fmt.Println(a.area())
	fmt.Println(b.area())
}
