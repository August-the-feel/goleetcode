package main

import (
	"fmt"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
// 输入：x = 121
// 输出：true

// 解题思路 进行取余数计算
// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
func isPalindrome(x int) bool {
	if (x <= 0) || (x%10 == 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	fmt.Println(isPalindrome(123))
}
