package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("aasfaasdfawer"))
}
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
