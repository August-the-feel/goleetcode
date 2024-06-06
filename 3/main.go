package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func main() {
	s := "asdzassdfr"
	res := lengthOfLongestSubstring_v2(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/*
* 暴力穷举法
* 时间复杂度: O(N^2)
* 空间复杂度: O(1)
 */
// 解题思路
// 进行遍历 不一样++ 一样的话进行删除 不一样end+1 一样的话start+1
func lengthOfLongestSubstring_v1(s string) int {
	ans := 0
	length := len(s)
	start, end := 0, 0
	// 一开始进行判断 不一样向前移动
	for end < length {
		for i := start; i < end; i++ {
			if s[i] == s[end] {
				start = i + 1
				break
			}
		}
		ans = max(ans, end-start+1)
		end++
	}
	return ans
}

/*
* 滑动窗口 + Hash
* Hash 表用于存储字母对应的最新下标
* 时间复杂度: O(N)
* 空间复杂度: O(1)
 */
// 解题思路
// 进行遍历 使用散列表 将不同的字符放到一个 map 中并且 如果存在将删除这之前的所有字符并计算长度
func lengthOfLongestSubstring_v2(s string) int {
	dict := make(map[byte]int)
	ans := 0
	left := -1
	for key, _ := range s {
		if val, ok := dict[s[key]]; ok {
			left = max(left, val)
		}
		dict[s[key]] = key
		ans = max(ans, key-left)
	}
	return ans
}

/*
* 动态规划 + Hash
* 设字符串为s，定义 dp[i] 为以 s[i] 为结尾的字符串的无重复字符的最长子串，
* 定义映射 lastindex map[byte][int] 记录在状态转移过程中，某字符出现的最末位置。
* 状态转移的三种情况:
* 1. s[i]没有在之前的状态出现过，即lastindex中找不到s[i], 则 dp[i] = dp[i-1] + 1
* 2. s[i]出现过，但是不包含在 dp[i-1] 对应的最大子串中， 表达式: dp[i-1] < i - lastindex[s[i]], 依旧是 dp[i] = dp[i-1] + 1
* 3. s[i]出现过，并且包含在 dp[i-1] 对应的最大子串中, 表达式: dp[i-1] >= i - lastindex[s[i]], 则: dp[i] = i - lastindex[s[i]]
 */
func lengthOfLongestSubstring_v3(s string) int {
	length := len(s)
	if length <= 0 {
		return 0
	} else if length == 1 {
		return 1
	}
	lastIndex := make(map[byte]int)
	dp := make([]int, length)
	ans := 0
	for i := 0; i < length; i++ {
		if index, ok := lastIndex[s[i]]; !ok || dp[i-1] < i-index {
			if i == 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-1] + 1
			}
		} else {
			dp[i] = i - index
		}
		lastIndex[s[i]] = i
		ans = max(ans, dp[i])
	}
	return ans
}

/*
* 动态规划 + Hash
* dp[i] 至于 dp[i-1] 有关，使用滚动的dp
* 设字符串为s，定义 dp[i] 为以 s[i] 为结尾的字符串的无重复字符的最长子串，
* 定义映射 lastindex map[byte][int] 记录在状态转移过程中，某字符出现的最末位置。
* 状态转移的三种情况:
* 1. s[i]没有在之前的状态出现过，即lastindex中找不到s[i], 则 dp[i] = dp[i-1] + 1
* 2. s[i]出现过，但是不包含在 dp[i-1] 对应的最大子串中， 表达式: dp[i-1] < i - lastindex[s[i]], 依旧是 dp[i] = dp[i-1] + 1
* 3. s[i]出现过，并且包含在 dp[i-1] 对应的最大子串中, 表达式: dp[i-1] >= i - lastindex[s[i]], 则: dp[i] = i - lastindex[s[i]]
 */
func lengthOfLongestSubstring_v4(s string) int {
	length := len(s)
	if length <= 0 {
		return 0
	} else if length == 1 {
		return 1
	}
	lastIndex := make(map[byte]int)
	ans := 0
	dp := 0
	for i := 0; i < length; i++ {
		idx, ok := lastIndex[s[i]]
		if !ok || dp < i-idx {
			dp = dp + 1
		} else {
			dp = i - idx
		}
		lastIndex[s[i]] = i
		ans = max(ans, dp)
	}
	return ans
}
