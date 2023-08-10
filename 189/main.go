package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	for _, v := range nums {
		fmt.Println(v)
	}

}

func rotate(nums []int, k int) {
	n := len(nums)
	a := make([]int, n)
	for i := 0; i < len(nums); i++ {
		// if j >= n {
		// 	// fmt.Println(int(math.Abs(float64(7 - j))))
		// 	// 第一版 负数出现错误
		// 	// a[int(math.Abs(float64(7-j)))] = nums[i]
		// 	// 优化第二版
		// 	a[int(j%n)] = nums[i]
		// } else {
		// 	a[i+k] = nums[i]
		// }

		// 优化第三版
		a[int((i+k)%n)] = nums[i]
	}
	copy(nums, a)
}
