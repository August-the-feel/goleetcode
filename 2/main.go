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
	list := &ListNode{Val: 2}
	list1 := &ListNode{Val: 6}
	list2 := &ListNode{Val: 4}
	list.Next = list1
	list1.Next = list2
	list2.Next = nil
	// for list != nil {
	// 	fmt.Println(list.Val)
	// 	list = list.Next
	// }
	root := &ListNode{Val: 2}
	root1 := &ListNode{Val: 4}
	root2 := &ListNode{Val: 3}
	root.Next = root1
	root1.Next = root2
	root2.Next = nil
	i := getIntersectionNode(list, root)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}
