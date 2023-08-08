package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list = &ListNode{}
	// 设置表头
	res := list
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			list.Next = list2
			list2 = list2.Next
		} else {
			list.Next = list1
			list1 = list1.Next
		}
		list = list.Next
	}
	if list1 != nil {
		list.Next = list1
	}
	if list2 != nil {
		list.Next = list2
	}
	return res.Next
}

func main() {
	lists := &ListNode{}
	lists.Val = 1
	lists.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}
	lists1 := &ListNode{}
	lists1.Val = 1
	lists1.Next = &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}

	i := mergeTwoLists(lists, lists1)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}
