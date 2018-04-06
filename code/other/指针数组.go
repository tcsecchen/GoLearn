package main

import (
	"fmt"
)

const max int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [max]*int

	for i = 0; i < max; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < max; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
