package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)
	for i, val := range nums {
		numMap[val] = i
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		complement := target - nums[i]
		if v, ok := numMap[complement]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{0, 0}
}

func main() {
	a := twoSum([]int{3, 2, 4, 8}, 6)
	fmt.Println(a)
}
