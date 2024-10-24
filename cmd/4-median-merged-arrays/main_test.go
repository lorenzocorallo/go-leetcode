package main

import (
	"fmt"
	"testing"
)

type input struct {
	nums1 []int
	nums2 []int
	exp   float64
}

var inputs = []input{
	{nums1: []int{1,2,3,4,5,6,7}, nums2: []int{-10, -9, -8, -7, -6}, exp: 1.5},
	{nums1: []int{1,2,3,4,5,6,7}, nums2: []int{6,7,8,9}, exp: 6.0},
	{nums1: []int{1, 2, 3, 4, 5, 6, 7}, nums2: []int{4, 5, 6, 7, 8, 9, 10, 11}, exp: 6.0},
	{nums1: []int{1, 3}, nums2: []int{2}, exp: 2.0},
	{nums1: []int{-10, 10}, nums2: []int{10}, exp: 10.0},
	{nums1: []int{-3, -2}, nums2: []int{-4, 1}, exp: -2.5},
	{nums1: []int{1, 1, 1, 1}, nums2: []int{}, exp: 1},
	{nums1: []int{1, 2, 3, 4}, nums2: []int{}, exp: 2.5},
	{nums1: []int{}, nums2: []int{1}, exp: 1},
	{nums1: []int{2}, nums2: []int{1}, exp: 1.5},
	{nums1: []int{1, 3}, nums2: []int{1, 3}, exp: 2.0},
	{nums1: []int{-1, 3}, nums2: []int{0, 2, 4}, exp: 2.0},
	{nums1: []int{1, 2, 3, 4, 5}, nums2: []int{2,3,4}, exp: 3.0},
	{nums1: []int{0,0,0,0,0}, nums2: []int{-1,0,0,0,0,0,1}, exp: 0.0},
}

func TestCustomSort(t *testing.T) {
	for _, in := range inputs {
		fmt.Printf("TEST CUSTOM SORT %+v, %+v\n", in.nums1, in.nums2)
		res := findMedianSortedArrays_CustomSort(in.nums1, in.nums2)
		if res != in.exp {
			t.Fatalf("exp: %f, res: %f\n", in.exp, res)
		}
	}
}

func Test(t *testing.T) {
	for _, in := range inputs {
		fmt.Printf("TEST DEFAULT SORT %+v, %+v\n", in.nums1, in.nums2)
		res := findMedianSortedArrays(in.nums1, in.nums2)
		if res != in.exp {
			t.Fatalf("exp: %f, res: %f\n", in.exp, res)
		}
	}
}
