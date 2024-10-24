package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, x1 := range nums[:] {
		for j, x2 := range nums[i+1:] {
			j = j + i + 1
			if x1+x2 == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Printf("%+v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%+v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%+v\n", twoSum([]int{3, 3}, 6))
}
