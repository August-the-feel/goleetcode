package main

import (
	"fmt"
	"math"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 示例 1：
//
//	输入：x = 123
//	输出：321
//
// 示例 2：
//
//	输入：x = -123
//	输出：-321
//
// 示例 3：
//
//	输入：x = 120
//	输出：21
//
// 示例 4：
//
//	输入：x = 0
//	输出：0

// 解题思路
// res*10 扩大10倍  x % 10 需要转换的字符串进行缩小10
func main() {
	s := 123
	res := reverse1(s)
	fmt.Println(res)
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		if int(int32(res)) != res {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	if int(int32(res)) != res {
		return 0
	}
	return res
}

func reverse1(x int) int {
	// 核心逻辑 s%10 然后进行
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + x%10 // res*10 扩大10倍  x % 10 需要转换的字符串进行缩小10
		x = x / 10
	}
	return res
}
