package main

import "fmt"

func twoSums(arr []int, sum int) (int, int) {
	numMap := make(map[int]int, len(arr))
	for i, val := range arr {
		if j, ok := numMap[sum-val]; ok {
			return j, i
		}
		numMap[val] = i
	}
	return -1, -1
}

func solveProblem(arr []int, sum int) {
	i, j := twoSums(arr, sum)
	fmt.Printf("array: %v, expect sum: %d, indexes: [%d, %d] --> number(%d, %d)",
		arr, sum, i, j, arr[i], arr[j])
}

func main() {
	solveProblem([]int{2, 7, 11, 15}, 9)
}
