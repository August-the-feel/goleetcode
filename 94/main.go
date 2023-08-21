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
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]

// Definition for a binary tree node.

func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	// 当二叉树的节点不为空或者栈中存储的数据不为空时
	for root != nil || len(stack) > 0 {
		// 如果节点不为空
		for root != nil {
			// 将节点指针存入数组
			stack = append(stack, root)
			// 将此节点换为节点的左子叶
			root = root.Left
		}
		// 如果节点为空，则返回上一个节点的指针
		root = stack[len(stack)-1]
		// 将上一个节点的指针进行出栈操作
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 将此节点换为节点的右子叶
		root = root.Right
	}
	return
}

// 使用递归方式处理
func inorderTraversal1(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 便利左边
		inorder(node.Left)
		res = append(res, node.Val)
		// 便利右边
		inorder(node.Right)
	}
	inorder(root)
	return
}
