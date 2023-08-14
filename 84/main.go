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
	// 创建一个单调栈 用于记录递增的柱子的索引。
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		// 如果当前柱子的高度小于等于栈顶元素所对应的柱子的高度，那么栈顶元素就不可能作为矩形的右边界了。因此，将栈顶元素所对应的柱子的右边界
		// 		设为当前柱子的索引， 并将栈顶元素出栈，直到栈顶元素的高度小于当前柱子的高度。
		// mono_stack[len(mono_stack)-1] 栈顶元素下标
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			// 满足 heights[mono_stack[len(mono_stack)-1]] >= heights[i] 栈顶的右边界 一直在 变化
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		// 左边的边界
		if len(mono_stack) == 0 {
			// 如果栈为空，则将当前柱左边界为 -1。
			left[i] = -1
		} else {
			// 如果栈不为空，那么栈顶元素所对应的柱子就是当前柱子的左边界。
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	fmt.Println(right)
	fmt.Println(left)

	// 创建一个最大值
	// ans := 0
	// for i := 0; i < n; i++ {
	// 	// (right[i]-left[i]-1)*heights[i] 求最大值 right[i]-left[i]-1 左边界 从-1 开始
	// 	ans = max(ans, (right[i]-left[i]-1)*heights[i])
	// }
	// return ans
	return 1
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	newHeights := make([]int, n+2)
	for i := 0; i < n; i++ {
		newHeights[i+1] = heights[i]
	}
	ret := 0               // 最大值
	left := make([]int, 0) // 创建 一个左边界
	for i, v := range newHeights {
		for len(left) > 0 && newHeights[left[len(left)-1]] > v {
			right := newHeights[left[len(left)-1]]
			left = left[:len(left)-1]
			ret = max(ret, right*(i-left[len(left)-1]-1))
		}
		left = append(left, i)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
