package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	fmt.Println(isSymmetric1(root))
}

// 给你一个二叉树的根节点 root ， 检查它是否轴对称
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
func isSymmetric(root *TreeNode) bool {
	return isSymmetricBtree(root, root)
}

func isSymmetricBtree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricBtree(p.Left, q.Right) && isSymmetricBtree(p.Right, q.Left)
}

func isSymmetric1(root *TreeNode) bool {
	i, j := root, root
	q := []*TreeNode{}
	q = append(q, i)
	q = append(q, j)
	for len(q) > 0 {
		i, j := q[0], q[1]
		q = q[2:]
		if i == nil && j == nil {
			continue
		}
		if i.Val != j.Val {
			return false
		}
		q = append(q, i.Left)
		q = append(q, j.Right)
		q = append(q, i.Right)
		q = append(q, j.Left)
	}
	return true
}
