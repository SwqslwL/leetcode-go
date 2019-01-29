package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	if head==nil{
		return nil
	}
	for ;p.Next != nil;{
		if p.Next.Val == p.Val{
			p.Next = p.Next.Next
		}else{
			p = p.Next
		}
	}
	return head
}


func main(){


}
