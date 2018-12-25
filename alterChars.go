package main

import "fmt"

// Complete the alternatingCharacters function below.
func alternatingCharacters(s string) int32 {
	count := int32(0)
	last := s[0]
	for i := 1; i < len(s); i++ {
		if last == s[i] {
			count++
		}
		last = s[i]
	}

	return count
}

func main() {
	var input string
	var output int32
	input = "AAAA"
	output = alternatingCharacters(input)
	fmt.Println(input, output)
	input = "BBBBB"
	output = alternatingCharacters(input)
	fmt.Println(input, output)
	input = "ABABABAB"
	output = alternatingCharacters(input)
	fmt.Println(input, output)
	input = "BABABA"
	output = alternatingCharacters(input)
	fmt.Println(input, output)
	input = "AAABBB"
	output = alternatingCharacters(input)
	fmt.Println(input, output)

}
