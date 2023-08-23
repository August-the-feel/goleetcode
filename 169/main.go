package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(nums))
}

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 示例 1：
// 输入：nums = [3,2,3]
// 输出：3

// 示例 2：
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
func majorityElement1(nums []int) int {
	// res 记录当前的元素 如果不一样就删除  一样就count++
	// 到最后只会有一个 大于众数的数字
	count, res := 0, 0
	for _, v := range nums {

		if count == 0 {
			res = v
		}
		if res == v {
			count++
		} else {
			count--
		}
	}
	return res
}
