package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	lists := head
	for lists.Next != nil {
		if lists.Next.Val == lists.Val {
			lists.Next = lists.Next.Next
		} else {
			lists = lists.Next
		}
	}
	return head
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
	i := deleteDuplicates(list)
	// deleteDuplicates(list)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}
