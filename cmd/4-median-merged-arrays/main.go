package main

import (
	"fmt"
	"slices"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	nums = slices.Concat(nums1, nums2)

	slices.Sort(nums)

	if len(nums)%2 == 1 {
		return float64(nums[(len(nums)-1)/2])
	} else {
		x1 := float64(nums[len(nums)/2])
		x2 := float64(nums[(len(nums)/2)-1])
		m := (x1+x2)/2
		return m
	}
}

func main() {
	fmt.Println("vim-go")
}
