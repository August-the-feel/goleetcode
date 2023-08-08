package main

import (
	"fmt"
)

func main() {
	num := []int{3, 2, 1, 1, 1, 123, 25, 4}
	fmt.Println(removeElement(num, 123))
}

func removeElement(nums []int, val int) int {
	result := []int{}
	for _, i := range nums {
		if i != val {
			result = append(result, i)
		}
	}
	return len(result)
}
