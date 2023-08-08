package main

import (
	"fmt"
)

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
