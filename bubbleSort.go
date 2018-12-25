package main

import (
	"fmt"
)

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	swaps := 0
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {

			if a[j] > a[j+1] {
				swaps++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.", swaps)
	fmt.Println()
	fmt.Println("First Element: ", a[0])
	fmt.Println("Last Element: ", a[n-1])
}

func main() {
	input := []int32{3, 2, 1}
	countSwaps(input)
	input = []int32{1, 2, 3}
	countSwaps(input)
}
