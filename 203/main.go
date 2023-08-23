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
	list := &ListNode{Val: 1}
	list1 := &ListNode{Val: 2}
	list2 := &ListNode{Val: 6}
	list3 := &ListNode{Val: 3}
	list5 := &ListNode{Val: 4}
	list6 := &ListNode{Val: 5}
	list7 := &ListNode{Val: 6}
	list.Next = list1
	list1.Next = list2
	list2.Next = list3
	list3.Next = list5
	list5.Next = list6
	list6.Next = list7
	list7.Next = nil
	// for list != nil {
	// 	fmt.Println(list.Val)
	// 	list = list.Next
	// }
	i := removeElements(list, 6)
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
	tmp := &ListNode{Next: head}
	for lists := tmp; lists.Next != nil; {
		if lists.Next.Val == val {
			lists.Next = lists.Next.Next
		} else {
			lists = lists.Next
		}
	}
	return tmp.Next
}
