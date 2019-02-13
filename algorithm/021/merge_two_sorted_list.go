package main

import (
	"fmt"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//add head
	head := new(ListNode)
	head.Next = l2
	p := l1
	q := head
	//traversal l2
	for p != nil {
		fmt.Println(q.Next)
		for q.Next != nil && p.Val > q.Next.Val {
			q = q.Next
		}
		node := p
		p = p.Next
		insertNode(q, node)
		q = node
	}
	return head.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1

	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

func insertNode(node *ListNode, inode *ListNode) {
	inode.Next = node.Next
	node.Next = inode
}

func main() {
	n1 := new(*ListNode)

	fmt.Printf("%T", n1)
}
