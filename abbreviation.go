package main

import (
	"fmt"
)

// Complete the abbreviation function below.
func abbreviation(a string, b string) string {

	n := len(a)
	m := len(b)
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, m+1)
	}
	d[0][0] = 1

	count := 0

	for k := 1; k <= n; k++ {
		i := k - 1
		if a[i] >= 65 && a[i] <= 90 || count == 1 {
			count = 1
			d[k][0] = 0
		} else {
			d[k][0] = 1
		}
	}

	for k := 1; k <= n; k++ {
		i := k - 1
		for l := 1; l <= m; l++ {
			j := l - 1
			if a[i] == b[j] {
				d[k][l] = d[k-1][l-1]
				continue
			} else {
				upper := a[i] ^ 0x20
				if upper == b[j] {
					d[k][l] = d[k-1][l-1] | d[k-1][l]
					continue
				}
			}
			if a[i] >= 65 && a[i] <= 90 {
				d[k][l] = 0
				continue
			} else {
				d[k][l] = d[k-1][l]
				continue
			}
		}
	}

	if d[n][m] > 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	var a, b, output string
	a = "daBcd"
	b = "ABC"
	output = abbreviation(a, b)
	fmt.Println(output)
	a = "AbcDE"
	b = "AFDE"
	output = abbreviation(a, b)
	fmt.Println(output)
}
