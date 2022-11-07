package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func atoi(s string) int {
	i := 0
	for _, c := range s {
		if c == ' ' {
			i++
		}
	}

	s = s[i:]

	sign := 1
	if strings.HasPrefix(s, "-") {
		sign = -1
		s = s[1:]
	} else if strings.HasPrefix(s, "+") {
		sign = 1
		s = s[1:]
	}

	result := 0

	for _, digit := range s {
		if unicode.IsDigit(digit) {
			result = result*10 + (int(digit) - int('0'))
		} else {
			break
		}

		if result*sign >= math.MaxInt32 {
			return math.MaxInt32
		} else if result*sign <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return result * sign
}

func solveProblem(s string) {
	fmt.Printf("string: %s, number: %d\n", s, atoi(s))
}

func main() {
	solveProblem("")
	solveProblem("42")
	solveProblem("   -42")
	solveProblem("4193 with words")
	solveProblem("words and 987")
	solveProblem("-91283472332")
}
