package main

import (
	"fmt"
	"sort"
)

type foo []int32

func (a foo) Len() int           { return len(a) }
func (a foo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a foo) Less(i, j int) bool { return a[i] < a[j] }

func abs(val int32) int32 {
	if val < 0 {
		return -1 * val
	}
	return val
}

func minimumAbsoluteDifference(arr []int32) int32 {
	sort.Sort(foo(arr))
	n := len(arr)
	minVal := abs(arr[0] - arr[1])

	for i := 1; i < n-1; i++ {
		val := abs(arr[i] - arr[i+1])

		if val < minVal {
			minVal = val
		}

	}

	return minVal
}

func main() {
	var input []int32
	var output int32
	input = []int32{3, -7, 0}
	output = minimumAbsoluteDifference(input)
	fmt.Println(output)
	input = []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}
	output = minimumAbsoluteDifference(input)
	fmt.Println(output)
	input = []int32{1, -3, 71, 68, 17}
	output = minimumAbsoluteDifference(input)
	fmt.Println(output)

}
