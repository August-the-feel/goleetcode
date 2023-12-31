package main

import (
	"fmt"
)

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]
// 解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//
//	你不需要考虑数组中超出新长度后面的元素。
//	例如，函数返回的新长度为 2 ，而
//	nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

func main() {
	num := []int{3, 2, 2, 3}
	fmt.Println(removeElement(num, 3))
}

func removeElement(nums []int, val int) int {
	// 使用切片的方式删除数据
	// result := []int{}
	// for _, i := range nums {
	// 	if i != val {
	// 		result = append(result, i)
	// 	}
	// }
	// nums = result
	// fmt.Println(nums)
	// return len(result)
	// 使用双指针1
	// left := 0
	// for _, v := range nums {
	// 	if v != val {
	// 		nums[left] = v
	// 		left++
	// 	}
	// }
	// fmt.Println(nums)
	// return left
	// 使用双指针2
	left, right := 0, len(nums)
	for left < right {
		// 相等将最后一个数字向前转移
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	fmt.Println(nums)
	return left
}
