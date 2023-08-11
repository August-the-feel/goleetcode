package main

import (
	"fmt"
)

func main() {
	var adf []int = []int{1, 3, 9, 9}
	fmt.Println(plusOne(adf))
}

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
// 解释：输入数组表示数字 123。
// 最后一位为9 则改为0 并向前第一个不为9的数 +1 ，
// 如果数组都为9 则 第一位为1 位数 +1 并改为0
func plusOne(digits []int) []int {
	digitsLen := len(digits)
	for i := digitsLen - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	// 其他情况均为9
	digits = make([]int, digitsLen+1)
	digits[0] = 1
	return digits
}
