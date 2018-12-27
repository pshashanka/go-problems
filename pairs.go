package main

import (
	"fmt"
	"sort"
)

type i32 []int32

func (a i32) Len() int {
	return len(a)
}

func (a i32) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a i32) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
	var n int32
	sort.Sort(i32(arr))
	i := 0

	for j := 1; j < len(arr); {
		diff := arr[j] - arr[i]
		if diff == k {
			n++
			j++
		}
		if diff < k {
			j++
		}
		if diff > k {
			i++
		}
	}

	return n
}

func main() {
	var input []int32
	var output, k int32
	input = []int32{1, 2, 3, 4}
	k = 1
	fmt.Println(input, k)
	output = pairs(k, input)
	fmt.Println(output)
	input = []int32{1, 5, 3, 4, 2}
	k = 2
	fmt.Println(input, k)
	output = pairs(k, input)
	fmt.Println(output)
}
