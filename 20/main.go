package main

import (
	"fmt"
)

func isValid(s string) bool {
	bytes := map[byte]byte{')': '(', ']': '[', '}': '{'}
	start := make([]byte, 0)
	if s == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		// 是 左边追加切片 不是就删除
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			start = append(start, s[i])
			// start[len(start)-1] == bytes[s[i]] 访问栈顶元素，即获取栈 start 中最后一个入栈的元素
		} else if len(start) > 0 && start[len(start)-1] == bytes[s[i]] {
			start = start[:len(start)-1]
		} else {
			return false
		}
	}
	return len(start) == 0
}

func main() {
	fmt.Println(isValid("{}()[]"))
}
