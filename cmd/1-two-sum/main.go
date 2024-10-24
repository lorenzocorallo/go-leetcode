package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, x := range nums {
		if j, ok := m[x]; ok {
			return []int{j,i}
		}

		diff := target - x
		m[diff] = i
		// (done on the upper if)
		// once on next iteration found x == diff,
		// return stored i and new i
		// the sum of the two nums is target 
	}

	return []int{0,0}
}

func main() {
	fmt.Printf("%+v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%+v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%+v\n", twoSum([]int{3, 3}, 6))
}
