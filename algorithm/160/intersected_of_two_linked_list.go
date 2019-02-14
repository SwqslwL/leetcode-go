package _60

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	head := new(ListNode)
	for p != nil && q != nil {
		p = p.Next
		q = q.Next
	}
	if p == nil {
		head = headB
		for q != nil {
			 q = q.Next
			 head = head.Next
		}
		q = headA
		for q != nil && head != nil {
			fmt.Println(head, q)
			if q == head {
				return head
			}
			q = q.Next
			head = head.Next
		}
	}else{
		head = headA
		for p != nil{
			p = p.Next
			head = head.Next
		}
		p = headB
		for p != nil && head != nil{
			fmt.Println(head)
			if p == head {
				return head
			}
			p = p.Next
			head = head.Next
		}
	}

	return nil
}


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	if p == nil || q == nil{
		return nil
	}
	for p != q{
		if p == nil {
			p = headB
		}else{
			p = p.Next
		}
		if q == nil {
			q = headA
		}else{
			q = q.Next
		}
	}
	return p
}