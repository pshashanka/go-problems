package main

import (
	"fmt"
	"sort"
)

type foo []int32

func (a foo) Len() int           { return len(a) }
func (a foo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a foo) Less(i, j int) bool { return a[i] < a[j] }

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	sort.Sort(foo(prices))
	var maxToys int32
	n := int32(len(prices))
	for maxToys = int32(0); maxToys < n && k > 0; maxToys++ {
		k -= prices[maxToys]
		fmt.Println(k)
	}
	return maxToys - 1
}

func main() {
	input := []int32{1, 12, 5, 111, 200, 1000, 10}
	k := int32(50)
	output := maximumToys(input, k)
	fmt.Println(output)
}
