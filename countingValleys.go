package main

import "fmt"

func countingValleys(n int32, s string) int32 {

	valleys := int32(0)
	vc := int32(0)
	for i := int32(0); i < n; i++ {

		if string(s[i]) == "D" {
			vc = vc - 1
		}

		if string(s[i]) == "U" {
			vc = vc + 1

			if vc == 0 {
				valleys = valleys + 1
			}
		}

	}

	return valleys
}

func main() {
	valleys := countingValleys(12, "DDUUDDUDUUUD")
	fmt.Println(valleys)

}
