package main

import (
	"fmt"
)

func main() {
	res := missingInteger([]int{3, 4, 5, 1, 12, 14, 13})
	fmt.Println(res)
}

func missingInteger(nums []int) int {
	sum := nums[0]
	dict := make(map[int]struct{})
	for i := 1; i < len(nums); i++ {
		if i < len(nums) && nums[i] != nums[i-1]+1 {
			break
		}
		sum += nums[i]
	}
	for _, val := range nums {
		dict[val] = struct{}{}
	}

	for range nums {
		if _, ok := dict[sum]; ok {
			sum++
		}
	}

	return sum
}
