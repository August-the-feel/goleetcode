package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(a))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	// 创建一个堆
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		// mono_stack
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			fmt.Println(heights[mono_stack[len(mono_stack)-1]])
			// 右边界
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
		fmt.Println(mono_stack)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
