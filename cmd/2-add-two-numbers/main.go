package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var last *ListNode = nil
	adder := 0

	for {
		fmt.Printf("l1 %+v\nl2 %+v\n", l1, l2)

		if l1 == nil && l2 == nil {
			break
		}

		val := adder
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		if val >= 10 {
			adder = val / 10
			val %= 10
		} else {
			adder = 0
		}

		last = &ListNode{
			Val:  val,
			Next: last,
		}
	}

	if adder >= 1 {
		last = &ListNode{
			Val:  adder,
			Next: last,
		}
	}

	var outlast *ListNode
	for i := 0; ; i *= 10 {
		if last == nil {
			break
		}
		outlast = &ListNode{
			Val:  last.Val,
			Next: outlast,
		}

		last = last.Next
	}

	return outlast
}
