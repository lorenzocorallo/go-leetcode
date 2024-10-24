package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	m := make(map[string]int)

	m[""] = 0
	m[" "] = 1
	m["b"] = 1
	m["abcabcbb"] = 3
	m["bbbbb"] = 1
	m["pwwkew"] = 3
	m["abcdefg"] = 7
	m["abcd fg"] = 7
	m["ab-d fg"] = 7
	m["ab-dbfg"] = 5

	for s, exp := range m {
		fmt.Printf("\nTEST \"%s\"\n", s)
		res := lengthOfLongestSubstring(s)
		if res != exp {
			t.Fatalf("str: %s, exp: %d, res: %d\n", s, exp, res)
		}
	}
}
