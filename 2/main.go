package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func main() {
	list := &ListNode{Val: 9}
	list1 := &ListNode{Val: 9}
	list2 := &ListNode{Val: 9}
	list3 := &ListNode{Val: 9}
	list.Next = list1
	list1.Next = list2
	list2.Next = list3
	// for list != nil {
	// 	fmt.Println(list.Val)
	// 	list = list.Next
	// }
	root := &ListNode{Val: 9}
	root1 := &ListNode{Val: 9}
	root2 := &ListNode{Val: 9}
	root3 := &ListNode{Val: 9}
	root4 := &ListNode{Val: 9}
	root5 := &ListNode{Val: 9}
	root.Next = root1
	root1.Next = root2
	root2.Next = root3
	root3.Next = root4
	root4.Next = root5
	root5.Next = nil
	i := addTwoNumbers(list, root)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 想完成两链表相加
	result := &ListNode{}
	dummy := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		dummy.Next = &ListNode{Val: sum % 10}
		dummy = dummy.Next
		carry = sum / 10
	}
	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
	}
	return result.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	ptr1 := l1
	ptr2 := l2
	cur := &ListNode{
		Val:  0,
		Next: nil,
	}
	ret := cur
	for {
		v1, v2 := 0, 0
		if ptr1 != nil {
			v1 = ptr1.Val
		}
		if ptr2 != nil {
			v2 = ptr2.Val
		}
		cur.Val = (v1 + v2 + add) % 10
		add = (v1 + v2 + add) / 10
		if ptr1 != nil {
			ptr1 = ptr1.Next
		}
		if ptr2 != nil {
			ptr2 = ptr2.Next
		}
		if ptr1 == nil && ptr2 == nil {
			if add > 0 {
				cur.Next = &ListNode{
					Val:  add,
					Next: nil,
				}
			}
			break
		} else {
			cur.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			cur = cur.Next
		}
	}
	return ret
}
