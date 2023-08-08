package main

import (
	"fmt"
)

func main() {
	fmt.Println(tribonacci(10))
}

// 泰波那契序列 Tn 定义如下：
// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
// 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
// 输入：n = 4
// 输出：4
// 解释：
// T_3 = 0 + 1 + 1 = 2
// T_4 = 1 + 1 + 2 = 4
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	q, w, e, r := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		q = w
		w = e
		e = r
		r = q + w + e
	}
	return r
}
