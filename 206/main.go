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
	list2 := &ListNode{Val: 3}
	list3 := &ListNode{Val: 4}
	list5 := &ListNode{Val: 5}
	list6 := &ListNode{Val: 6}
	list7 := &ListNode{Val: 7}
	list.Next = list1
	list1.Next = list2
	list2.Next = list3
	list3.Next = list5
	list5.Next = list6
	list6.Next = list7
	list7.Next = nil
	i := reverseList(list)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}

func reverseList(head *ListNode) *ListNode {

}
