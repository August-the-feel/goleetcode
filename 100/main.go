package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{}
	root.Val = 1
	root1 := &TreeNode{}
	root1.Val = 3
	root2 := &TreeNode{}
	root2.Val = 2
	root.Left = nil
	root.Right = root1
	root1.Right = root2

	too := &TreeNode{}
	too.Val = 1
	too.Left = nil
	too.Right = root1
	fmt.Println(isSameTree(root, too))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
