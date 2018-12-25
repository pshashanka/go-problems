package main

import "fmt"

func slice(row int32, col int32, arr [][]int32) int32 {
	var val int32
	val += arr[row][col] + arr[row][col+1] + arr[row][col+2]
	val += arr[row+1][col+1]
	val += arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
	fmt.Println(row, col, val)
	return val
}

func hourglassSum(arr [][]int32) int32 {
	var max int32

	for row := int32(0); row < int32(len(arr)-2); row++ {
		for col := int32(0); col < int32(len(arr[row])-2); col++ {
			sliceVal := slice(row, col, arr)

			if row == 0 && col == 0 {
				max = sliceVal
			} else if max < sliceVal {
				max = sliceVal
			}
		}

	}

	return max
}

func main() {
	input :=
		[][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
	var ma int32
	ma = hourglassSum(input)
	fmt.Println(ma)
}
