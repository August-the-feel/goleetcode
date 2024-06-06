package main

import (
	"fmt"
)

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

// 解题死思路
// 将需要的数字添加到一个新的数组中 寻找，找到返回没有找到返回-1
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
	fmt.Println(a)
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
