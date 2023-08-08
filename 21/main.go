package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将两个升序链表合并为一个新的 升序 链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
// 链表1：1->2->3
// 链表2：1->3->4
// 合链表：1->1->2->3->3->4
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
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
