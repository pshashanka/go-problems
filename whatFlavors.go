package main

import "fmt"

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
	mp := make(map[int32]int32)

	for pos, k := range cost {
		diff := money - k
		in, found := mp[diff]
		if found {
			fmt.Printf("%d %d", in+1, pos+1)
			fmt.Println()
		} else if diff > 0 && diff < money {
			mp[k] = int32(pos)
		}
	}

}

func main() {
	var input []int32
	var money int32
	input = []int32{1, 4, 5, 3, 2}
	money = 4
	whatFlavors(input, money)
	input = []int32{2, 2, 4, 3}
	money = 4
	whatFlavors(input, money)
}
