package main

import (
	"fmt"
)

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

// 字母和数字都属于字母数字字符。

// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	var a map[int]int
	a = make(map[int]int, 10)
	for i, x := range nums {
		if p, ok := a[target-x]; ok {
			return []int{i, p}
		}
		a[x] = i
	}
	return nil
}

func main() {
	var a []int = []int{2, 7, 11, 15, 10}
	fmt.Println(twoSum1(a, 9))
}