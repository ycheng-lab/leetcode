package main

import "fmt"

func calcMedian(one, two []int) float32 {
	whole := merge(one, two)
	return findMedian(whole)
}

func merge(one, two []int) []int {
	whole := make([]int, len(one)+len(two))
	i, j := 0, 0

	for index := 0; index < (len(one) + len(two)); index++ {
		if j == len(two) || (i < len(one) && j < len(two) && one[i] < two[j]) {
			whole[index] = one[i]
			i++
			continue
		}

		if i == len(one) || (i < len(one) && j < len(two) && one[i] >= two[j]) {
			whole[index] = two[j]
			j++
		}
	}

	return whole
}

func findMedian(arr []int) float32 {
	k := len(arr) / 2
	if len(arr)%2 == 1 {
		return float32(arr[k])
	} else {
		return float32(arr[k]+arr[k-1]) / 2.0
	}
}

func solveProblem(one, two []int) {
	whole := merge(one, two)
	median := calcMedian(one, two)
	fmt.Printf("array one: %v\n", one)
	fmt.Printf("array two: %v\n", two)
	fmt.Printf("array whole: %v\n", whole)
	fmt.Printf("median: %f\n", median)

}

func main() {
	solveProblem([]int{2, 4}, []int{7, 8, 9})

	solveProblem([]int{2, 4, 8}, []int{7, 8, 9})
}
