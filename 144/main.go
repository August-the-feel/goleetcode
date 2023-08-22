package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ins func(node *TreeNode)

func main() {
	// 前序遍历是 当前节点 左节点 右节点
	root := &TreeNode{}
	root.Val = 1
	root1 := &TreeNode{}
	root1.Val = 3
	root2 := &TreeNode{}
	root2.Val = 2
	root.Left = nil
	root.Right = root1
	root1.Right = root2
	fmt.Println(preorderTraversal(root))
}

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]

// 前序遍历是 当前节点 左节点 右节点

func preorderTraversal(root *TreeNode) (res []int) {
	ins = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		ins(node.Left)
		ins(node.Right)
	}
	ins(root)
	return
}
