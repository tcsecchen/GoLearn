package main

import (
	"fmt"
)

func main() {
	var i,j int64;
	for i = 1 ;i <= 9 ;i++{
		for j = 1;j <= i; j++{
			fmt.Printf("%d*%d=%d ", i, j, i*j);
		}
		fmt.Printf("\n");
	}
}