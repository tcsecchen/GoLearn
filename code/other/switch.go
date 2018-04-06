package main

import (
	"fmt"
	"os"
)

func mySwitch() {
	a, b, c := 0, 0, 0
	for _, opcase := range os.Args[1:] {
		switch opcase {
		case "1":
			a = 1
			fmt.Println("1")
		case "2":
			b = 2
			fmt.Println("2")
		case "3":
			c = 3
			fmt.Println("3")
		default:
			a, b, c = 0, 0, 0
			fmt.Println("default")
		}
	}
	println(a, b, c)
	//交换
	a, b = b, a
	println(a, b, c)
}
