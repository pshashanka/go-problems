package main

import (
	"fmt"
)

func minimumSwaps(arr []int32) int32 {
	fmt.Println(arr)
	var swaps int32
	var n = len(arr)
	for i := 0; i+1 < n; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
				break
			}
		}

		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
			swaps++
		}

	}
	// fmt.Println(arr)
	return swaps
}

func main() {
	var input []int32
	input = []int32{4, 3, 1, 2}
	fmt.Println(minimumSwaps(input), 3)
	input = []int32{2, 3, 4, 1, 5}
	fmt.Println(minimumSwaps(input), 3)
	input = []int32{1, 3, 5, 2, 4, 6, 7}
	fmt.Println(minimumSwaps(input), 3)
}
