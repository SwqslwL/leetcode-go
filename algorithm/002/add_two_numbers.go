package main

import (
	"fmt"

)

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	var carry = 0 //jinwei
	var sum = 0

	head  := new(ListNode)
	move := head

	for ;l1 != nil||l2 != nil||carry!=0;{
		//fmt.Println(l1.Val,l2.Val)
		node := new(ListNode)//添加节点
		var m,n int
		if l1 == nil{
			m =0
		}else{
			m = l1.Val
			l1 = l1.Next
		}
		if l2 == nil{
			n =0
		}else{
			n=l2.Val
			l2 = l2.Next
		}
		sum = m + n + carry
		if sum >9{//need
			node.Val = sum - 10
			carry = 1
		}else{
			node.Val = sum
			carry = 0
		}
		//fmt.Println(node.Val)
		move.Next = node
		move = move.Next
	}
	head = head.Next
	getNum(head)
	return head
}

func getNum(l1 *ListNode){
	var head = l1
	for ;head != nil;head=head.Next{
		fmt.Print(head.Val)
	}
}

func createListNode(num1 string) *ListNode{
	var rear *ListNode = nil
	for i:=0;i<len(num1);i++{
		node := ListNode{Val:int(num1[i]-48),Next:rear}
		rear = &node
	}
	return rear
}

func main(){
	num1 := string("10000000000000000000000000000000001")
	num2 := string("4")
	l1 := createListNode(num1)
	l2 := createListNode(num2)
	addTwoNumbers(l1,l2)
	//fmt.Println(x)
}