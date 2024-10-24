package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 1
	lastIndex := make([]int , 128)
	startIdx := 0

	for i, c := range s {
		subIdxStart := max(startIdx, lastIndex[c])
		subIdxEnd := i
		subLen := 1 + subIdxEnd - subIdxStart
		startIdx = max(startIdx, subIdxStart)
		maxLen = max(maxLen, subLen)
		lastIndex[c] = i + 1
	}

	return maxLen
}

func main() {
	fmt.Println("vim-go")
}
