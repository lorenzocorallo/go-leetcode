package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func makeList(nums []int) *ListNode {
	slices.Reverse(nums)
	var out *ListNode
	for _, v := range nums {
		out = &ListNode{
			Val:  v,
			Next: out,
		}
	}

	return out
}

func makeSlice(list *ListNode) []int {
	out := make([]int, 0)
	tmp := list
	for {
		out = append(out, tmp.Val)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

	return out
}

func Test1(t *testing.T) {
	first := makeList([]int{2, 4, 3})
	second := makeList([]int{5, 6, 4})

	exp := makeList([]int{7, 0, 8})
	res := addTwoNumbers(first, second)

	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("exp: %+v\nres:%+v\n", makeSlice(exp), makeSlice(res))
	}
}

func Test2(t *testing.T) {
	fmt.Print("\nTEST\n")
	first := makeList([]int{9, 9, 9, 9, 9, 9, 9})
	second := makeList([]int{9, 9, 9, 9})

	exp := makeList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	res := addTwoNumbers(first, second)
	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("exp: %+v\nres:%+v\n", makeSlice(exp), makeSlice(res))
	}
}
