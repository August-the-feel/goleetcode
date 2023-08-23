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
	root := &ListNode{Val: 1}
	root1 := &ListNode{Val: 2}
	root2 := &ListNode{Val: 6}
	root3 := &ListNode{Val: 3}
	root5 := &ListNode{Val: 4}
	root6 := &ListNode{Val: 5}
	root7 := &ListNode{Val: 6}
	root.Next = root1
	root1.Next = root2
	root2.Next = root3
	root3.Next = root5
	root5.Next = root6
	root6.Next = root7
	root7.Next = list5
	// for root != nil {
	// 	fmt.Println(root.Val)
	// 	root = root.Next
	// }
	i := getIntersectionNode(list, root)
	// deleteDuplicates(list)
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}

// 给你两个单链表的头节点 headA 和 headB ,
//
//	请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
//
// 图示两个链表在节点 c1 开始相交：
// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。

// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'
// 解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
// 在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
// — 请注意相交节点的值不为 1，因为在链表 A 和链表 B 之中值为 1 的节点 (A 中第二个节点和 B 中第三个节点) 是不同的节点。
//
//	换句话说，它们在内存中指向两个不同的位置，而链表 A 和链表 B 中值为 8 的节点 (A 中第三个节点，B 中第四个节点) 在内存中指向相同的位置。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stack := map[*ListNode]bool{}
	for lists := headA; lists != nil; lists = lists.Next {
		stack[lists] = true
	}
	for lists := headB; lists != nil; lists = lists.Next {
		if stack[lists] {
			return lists
		}
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	res1, res2 := headA, headB
	for res1 != res2 {
		if res1 == nil {
			res1 = headB
		} else {
			res1 = res1.Next
		}
		if res2 == nil {
			res2 = headA
		} else {
			res2 = res2.Next
		}
	}
	return res1
}
