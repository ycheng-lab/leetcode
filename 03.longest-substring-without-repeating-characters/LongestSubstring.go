package main

import "fmt"

func longestSubstring(s string) (string, int) {

	locations := make([]int, 256)
	for i := range locations {
		locations[i] = -1
	}

	index, maxLength, left := 0, 0, 0
	for i, value := range s {
		if locations[value] >= left {
			left = locations[value] + 1
		} else if i-left+1 > maxLength {
			index = i
			maxLength = index - left + 1
		}
		locations[value] = i
	}

	return s[index-maxLength+1 : index+1], maxLength
}

func solveProblem(s string) {
	substr, length := longestSubstring(s)
	fmt.Printf("string: (%s), substr: (%s), length=%d\n", s, substr, length)
}

func main() {
	solveProblem("abcabcbb")
	solveProblem("bbbbb")
	solveProblem("pwwkew")
}
