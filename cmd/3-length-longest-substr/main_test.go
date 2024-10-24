package main

import (
	"fmt"
	"testing"
)

type input struct {
	s   string
	exp int
}

func Test(t *testing.T) {
	inputs := []input{
		{s: "", exp: 0},
		{s: " ", exp: 1},
		{s: "b", exp: 1},
		{s: "abcabcbb", exp: 3},
		{s: "bbbbb", exp: 1},
		{s: "pwwkew", exp: 3},
		{s: "abcdefg", exp: 7},
		{s: "abcd fg", exp: 7},
		{s: "ab-d fg", exp: 7},
		{s: "ab-dbfg", exp: 5},
		{s: "tmmzuxt", exp: 5},
	}

	for _, in := range inputs {
		fmt.Printf("\nTEST \"%s\"\n", in.s)
		res := lengthOfLongestSubstring(in.s)
		if res != in.exp {
			t.Fatalf("str: %s, exp: %d, res: %d\n", in.s, in.exp, res)
		}
	}
}
