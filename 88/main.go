package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums1 := []int{4, 5, 6}
	n := 3
	merge(nums, m, nums1, n)
	for _, v := range nums {
		fmt.Println(v)
	}
}

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
//
//	为了应对这种情况，nums1 的初始长度为 m + n，
//	其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。
//	nums2 的长度为 n 。
//	输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
func merge(nums1 []int, m int, nums2 []int, n int) {
	a := make([]int, 0, m+n)
	q, p := 0, 0
	for {
		if q == m {
			a = append(a, nums2[p:]...)
			break
		}
		if p == n {
			a = append(a, nums1[q:]...)
			break
		}
		if nums1[q] < nums2[p] {
			a = append(a, nums1[q])
			q++
		} else {
			a = append(a, nums2[p])
			p++
		}
	}
	for _, v := range a {
		fmt.Println(v)
	}
	copy(nums1, a)
}
