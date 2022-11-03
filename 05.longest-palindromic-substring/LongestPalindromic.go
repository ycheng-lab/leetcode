package main

import "fmt"

func longestPalindrome(s string) string {
	maxLen := 0
	maxPalindrome := ""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if isPalindrome(s, i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					maxPalindrome = s[i : j+1]
				}
			}
		}
	}
	return maxPalindrome
}

func isPalindrome(s string, i, j int) bool {
	if i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func solveProblem(s string) {
	fmt.Printf("string: %s, longest palindrome: %s\n", s, longestPalindrome(s))
}

func main() {
	solveProblem("babad")
	solveProblem("cbbd")
}
