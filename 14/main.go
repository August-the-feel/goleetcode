package main

import (
	"fmt"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs) || (strs[j][i] != strs[0][j]) {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}

func main() {
	var ss []string = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(ss))
}
