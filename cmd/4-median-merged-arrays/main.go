package main

import (
	"fmt"
	"slices"
)

func getMedian(nums []int) float64 {
	if len(nums)%2 == 1 {
		return float64(nums[(len(nums)-1)/2])
	} else {
		x1 := float64(nums[len(nums)/2])
		x2 := float64(nums[(len(nums)/2)-1])
		m := (x1+x2)/2
		return m
	}  
}

func findMedianSortedArrays_CustomSort(nums1, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMedian(nums2)
	}

	if len(nums2) == 0 {
		return getMedian(nums1)
	}

	nums := make([]int, len(nums1)+len(nums2))

	if nums1[0] >= nums2[len(nums2) - 1] {
		nums = slices.Concat(nums2, nums1)
		return getMedian(nums)
	}

	if nums2[0] >= nums1[len(nums1) - 1] {
		nums = slices.Concat(nums1, nums2)
		return getMedian(nums)
	}

	if len(nums1) == 1 || len(nums2) == 1 {
		nums = slices.Concat(nums1, nums2)
		slices.Sort(nums)
		return getMedian(nums)
	}


	var t1, t2 []int
	if nums1[0] <= nums2[0] {
		t1 = nums1
		t2 = nums2
	} else {
		t1 = nums2
		t2 = nums1
	}

	//fmt.Printf("t1 %+v -- t2 %+v\n", t1, t2)

	minIdx := -1
	maxIdx := -1
	for i := 0; i < len(t1); i++ {
		if t1[i] >= t2[0] {

			//fmt.Printf("FOUND MIN!!! t1[%d] = %d, t2[0] = %d\n", i, t1[i], t2[0])
			minIdx = i - 1
			break 
		}
	}

	for j := 0; j < len(t2); j++ {
		x := t1[len(t1)-1]
		if x <= t2[j] {
			//fmt.Printf("FOUND MAX!!! t1[last] = %d, t2[%d] = %d\n", x, j, t2[j])
			maxIdx = j + 1
			break 
		}
	}

	if maxIdx == -1 && t2[len(t2)-1] <= t1[len(t1)-1] {
		nums = slices.Concat(nums1, nums2)
		slices.Sort(nums)
		return getMedian(nums)
	}

	var s1, s2, s3 []int
	if maxIdx == -1 {
		s1 = t1[:minIdx+1]
		s2 = t2
		s3 = t1[minIdx+1:]
	} else {
		s1 = t1[:minIdx+1]
		s2 = slices.Concat(t1[minIdx+1:], t2[:maxIdx])
		slices.Sort(s2)
		s3 = t2[maxIdx:]
	}

	//fmt.Printf("%d, %d\n", minIdx, maxIdx)
	//fmt.Printf("%+v %+v %+v\n", s1, s2, s3)
	nums = slices.Concat(s1, s2, s3)
	//fmt.Printf("sorted nums %+v", nums)


	return getMedian(nums)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	nums = slices.Concat(nums1, nums2)

	slices.Sort(nums)

	if len(nums)%2 == 1 {
		return float64(nums[(len(nums)-1)/2])
	} else {
		x1 := float64(nums[len(nums)/2])
		x2 := float64(nums[(len(nums)/2)-1])
		m := (x1+x2)/2
		//fmt.Printf("x1 %f, x2 %f, m %f", x1, x2, m)
		return m
	}  
}

func main() {
	fmt.Println("vim-go")
}
