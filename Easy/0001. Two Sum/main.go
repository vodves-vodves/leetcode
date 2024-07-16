package main

import (
	"fmt"
)

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for i1, j1 := range nums {
			if j+j1 == target && i != i1 {
				return []int{i, i1}
			}
		}
	}
	return nil
}
