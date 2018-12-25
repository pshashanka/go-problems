package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	jumps := int32(0)
	zerNext := 0
	for i := 0; i < len(c); i++ {

		if c[i] == 1 && i < (len(c)-1) {
			jumps++
			zerNext = 0
		}

		if c[i] == 0 {
			zerNext++
			if zerNext == 2 {
				zerNext = 0
				jumps++
			}
		}

	}

	return jumps
}

func main() {
	jumps := jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0})
	fmt.Println(jumps)

	jumps = jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})
	fmt.Println(jumps)

	jumps = jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0})
	fmt.Println(jumps)
}
