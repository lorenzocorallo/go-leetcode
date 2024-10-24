package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 1
	for i, c1 := range s {
		set := make(map[rune]bool)
		set[c1] = true
		for j, c2 := range s[i+1:] {
			j += i + 1
			if _, ok := set[c2]; ok {
				maxLen = max(j-i, maxLen)
				break
			}

			set[c2] = true
			
			if j == len(s) - 1 {
				maxLen = max(j-i+1, maxLen)
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println("vim-go")
}
