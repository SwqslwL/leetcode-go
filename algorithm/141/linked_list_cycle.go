package _41

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p := head
	dict := make(map[*ListNode]int)
	if p == nil{
		return false
	}
	for {
		//将节点存储在map中
		dict[p] = p.Val
		if p.Next == nil{
			return false
		}else if _,ok := dict[p.Next];ok{
			return true
		}else{
			p = p.Next
		}
	}
}

//双指针法
func hasCycle2(head *ListNode) bool {
	slow := head
	fast := head
	if head == nil {
		 return false
	}
	for {
		for i:=0 ; i<=1 ;i++{
			if !run(&fast) {
				return false
			}
			if fast.Next == slow {
				return true
			}
		}
		run(&slow)
	}
}

func run(p **ListNode) bool {
	*p = (*p).Next
	if *p == nil {
		return false
	}
	return true
}

//
func hasCycle3(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}