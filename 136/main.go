package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int = []int{7, 7, 4, 4, 1}
	fmt.Println(singleNumber1(a))
}

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。
//	找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
// 示例 1 ：
// 输入：nums = [2,2,1]
// 输出：1
// 示例 2 ：
// 输入：nums = [4,1,2,1,2]
// 输出：4

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	newStack := make([]int, 0, len(nums))
	for _, v := range nums {
		if len(newStack) > 0 && newStack[len(newStack)-1] == v {
			newStack = newStack[:len(newStack)-1]
		} else {
			newStack = append(newStack, v)
		}
	}
	return newStack[0]
}

func singleNumber1(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
