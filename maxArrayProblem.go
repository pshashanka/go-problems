package main

import (
	"fmt"
)

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	l := len(arr)

	if l <= 0 {
		return 0
	}

	if l == 1 {
		return arr[0]
	}

	arr[1] = max(arr[0], arr[1])
	for i := 2; i < l; i++ {
		arr[i] = max(arr[i-1], max(arr[i], max(arr[i]+arr[i-2], arr[i-2])))
	}

	return arr[l-1]
}

func main() {

	var input []int32
	var output int32
	input = []int32{3, 7, 4, 6, 5}
	fmt.Print(input)
	output = maxSubsetSum(input)
	fmt.Println(output)
	input = []int32{2, 1, 5, 8, 4}
	fmt.Print(input)
	output = maxSubsetSum(input)
	fmt.Println(output)
	input = []int32{3, 5, -7, 8, 10}
	fmt.Print(input)
	output = maxSubsetSum(input)
	fmt.Println(output)
}
