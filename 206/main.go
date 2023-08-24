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
	// 创建一个第三变量

	// var prev *ListNode 声明一个指针变量，通常用于在遍历链表等情况下，
	// 	通过指针变量来跟踪前一个节点。可以在后续操作中，根据需求对 prev
	// 	进行赋值和操作。
	// res := &ListNode{} 声明并初始化一个指向 ListNode 结构体实例的指针变量。
	// 	通常用于创建一个空节点或作为结果节点的头节点。
	// 	在后续操作中，可以对 res 指针所指向的结构体实例进行成员属性的读写操作。

	var res *ListNode // 最后的链表
	stack := head     // 中间变量
	for stack != nil {
		next := stack.Next // 将下一个节点 存放到临时变量
		stack.Next = res   // 将下一个赋空值
		res = stack
		stack = next //
	}
	return res
}
