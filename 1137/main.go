package main

import (
	"fmt"
)

func main() {
	fmt.Println(tribonacci(10))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	q, w, e, r := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		q = w
		w = e
		e = r
		r = q + w + e
	}
	return r
}
