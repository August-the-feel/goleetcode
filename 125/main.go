package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("aasfaasdfawer"))
}

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var schar string
	for i := 0; i < len(s); i++ {
		if isachar(s[i]) {
			schar += string(s[i])
		}
	}
	// 翻转字符串
	// var sflipchar string
	// for i := len(schar) - 1; i >= 0; i-- {
	// 	sflipchar += string(schar[i])
	// }
	// if strings.EqualFold(sflipchar, schar) {
	// 	return true
	// } else {
	// 	return false
	// }

	// 使用双指针
	left, right := 0, len(schar)-1
	if left < right {
		if schar[left] != schar[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isachar(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}
