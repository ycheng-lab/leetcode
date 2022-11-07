package main

import "fmt"

func reverseInt(n int) int {
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	result := 0
	for n > 0 {
		result *= 10
		result += (n % 10)
		n = n / 10
	}

	return sign * result
}

func solveProblem(n int) {
	fmt.Printf("original: %d, reversed: %d\n", n, reverseInt(n))
}

func main() {
	solveProblem(8)
	solveProblem(123)
	solveProblem(-979823)
	solveProblem(-2)
	solveProblem(0)
}
