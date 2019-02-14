package main

import (
	"fmt"
	"leetcode/utils"
	_ "time"
)

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// without head
func createLinkedList(n int) *LinkedList {
	L := new(LinkedList)
	head := L
	for i := 0; i < n; i++ {
		node := LinkedList{Val: utils.GetRandNum(), Next: nil}
		L.Next = &node
		L = &node
	}
	return head.Next
}

func getLinkedList(L *LinkedList) {
	p := L
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

func main() {
	L := createLinkedList(4)
	getLinkedList(L)
}
