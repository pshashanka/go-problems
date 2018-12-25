package main

import (
	"fmt"
	"strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for _, b := range alpha {
		sub := string(b)
		if strings.Contains(s1, sub) && strings.Contains(s2, sub) {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var output string

	output = twoStrings("hello", "world")
	fmt.Println(output)
	output = twoStrings("hi", "world")
	fmt.Println(output)
}
