package main

import "fmt"

func isPalindromeNumber(n int) bool {
	if n < 0 {
		return false
	}

	for n > 0 {

		if n < 10 {
			return true
		}

		i := 1
		left := n / 10
		for left > 10 {
			left /= 10
			i++
		}

		if left != n%10 {
			return false
		}

		for ; i > 0; i-- {
			left *= 10
		}

		n -= left
		n /= 10
	}

	return true
}

func solveProblem(n int) {
	fmt.Printf("%d is palindrome number: %v\n", n, isPalindromeNumber(n))
}

func main() {
	solveProblem(2)
	solveProblem(28273642)
	solveProblem(123456654321)
	solveProblem(17361)
}
