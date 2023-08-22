package main

import (
	"fmt"
)

func main() {
	// 后序遍历 左节点 右节点 当前节点
	root := &TreeNode{}
	root.Val = 1
	root1 := &TreeNode{}
	root1.Val = 2
	root2 := &TreeNode{}
	root2.Val = 3
	root.Left = nil
	root.Right = root1
	root1.Right = root2
	fmt.Println(postorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ins func(node *TreeNode)

// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
// 后序遍历 左节点 右节点 当前节点

// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[3,2,1]
// 示例 2：
// 输入：root = []
// 输出：[]

func postorderTraversal(root *TreeNode) (res []int) {
	ins = func(node *TreeNode) {
		if node == nil {
			return
		}
		ins(node.Left)
		ins(node.Right)
		res = append(res, node.Val)

	}
	ins(root)
	return
}
