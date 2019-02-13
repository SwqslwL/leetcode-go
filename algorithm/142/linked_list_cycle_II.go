package _42

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	Node := make(map[*ListNode]int)
	p := head
	for p != nil {
		Node[p] = p.Val
		p = p.Next
		if _, ok := Node[p]; ok {
			return p
		}
	}
	return nil
}

//双指针
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	p := head
	q := slow
	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}
