package main

import (
	"fmt"
	"sort"
)

type ddArr [][]int32

func (a ddArr) Len() int      { return len(a) }
func (a ddArr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ddArr) Less(i, j int) bool {

	if a[i][1] != a[j][1] {
		return a[i][1] < a[j][1]
	}

	if a[i][1] == a[j][1] {
		return a[i][0] > a[j][0]
	}

	return false
}

// Complete the luckBalance function below.
func luckBalance(k int32, contests [][]int32) int32 {
	var max int32
	toLose := int32(0)
	sort.Sort(ddArr(contests))
	fmt.Println("Input: ", contests, k)
	for i := 0; i < len(contests); i++ {
		if contests[i][1] == 0 {
			max += contests[i][0]
		}
		if contests[i][1] == 1 && toLose >= k {
			max -= contests[i][0]
		}
		if contests[i][1] == 1 && toLose < k {
			max += contests[i][0]
			toLose++
		}

	}

	return max
}

func main() {
	var contests [][]int32
	var output int32
	var k int32
	contests = [][]int32{
		{5, 1},
		{2, 1},
		{1, 1},
		{8, 1},
		{10, 0},
		{5, 0},
	}
	k = 3
	fmt.Println("Input: ", contests, k)
	output = luckBalance(k, contests)
	fmt.Println("Output", output)
}
