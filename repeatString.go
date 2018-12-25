package main

import (
	"fmt"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	count := int64(strings.Count(s, "a"))
	l := int64(len(s))
	mul := int64(n / l)
	count = count * mul
	rem := n % l

	for i := int64(0); i < rem; i++ {
		if string(s[i]) == "a" {
			count++
		}

	}

	return count
}

func main() {
	var ans int64
	ans = repeatedString("aba", 10)
	fmt.Println(ans)
	ans = repeatedString("a", 1000000000000)
	fmt.Println(ans)

}
