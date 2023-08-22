package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		Val: 1,
	}
	list.Next = &ListNode{Val: 1}
	list.Next.Next = &ListNode{Val: 2}
	list.Next.Next.Next = nil
	// for list != nil {
	// 	fmt.Println(list.Val)
	// 	list = list.Next
	// }
	i := removeElements(list, 1)
	// deleteDuplicates(list)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
//	示例 1：
//
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
// 示例 2：
// 输入：head = [], val = 1
// 输出：[]
// 示例 3：
// 输入：head = [7,7,7,7], val = 7
// 输出：[]
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	lists := head
	for lists.Next != nil {
		if lists.Val == val {
			lists.Next = lists.Next.Next
		} else {
			lists = lists.Next
		}
	}
	return lists
}
