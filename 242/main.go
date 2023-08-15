package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram1("anagram", "nagaram"))
}

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true

// 使用map 解决问题
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 不能使用 map 比较
	num := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := num[s[i]]; !ok {
			num[s[i]] = 1
		} else {
			num[s[i]] = val + 1
		}
	}
	for i := 0; i < len(t); i++ {
		if val, ok := num[t[i]]; ok {
			val--
			num[t[i]] = val
			if num[t[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// 长度相同则可以一个循环出来
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for i := 0; i < len(t); i++ {
		c1[s[i]-'a']++
		c2[t[i]-'a']++
	}
	return c1 == c2
}
