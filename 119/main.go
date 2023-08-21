package main

import (
	"fmt"
)

func main() {
	fmt.Println(getRow(3))
}

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 输入: rowIndex = 3
// 输出: [1,3,3,1]

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+1)
	for i, _ := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res[rowIndex]
}
