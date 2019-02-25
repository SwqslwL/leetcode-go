package tree

import (
	"fmt"
	"leetcode/structure/queue"
)
type Tree interface {
	CreateTreeByLevel()
}

type BiTree struct {
	lchild *BiTree
	rchild *BiTree
	Val interface{}
}


//series =
func CreateTreeByLevel (series []interface{}) *BiTree{
	q := new(queue.Queue)
	q.InitQueue()
	node := new(BiTree)
	if len(series) != 0 {   //根结点入队
		node.Val = series[0]
		q.EnQueue(node)
	}
	for i := 1 ; i < len(series) ; i += 2 {
		node := new(BiTree)
		node = q.DeQueue().(*BiTree)
		lc := new(BiTree)
		lc.Val = series[i]
		if lc.Val != nil {
			node.lchild = lc
			q.EnQueue(node.lchild)
		}
		if i + 1 < len(series) {
			rc := new(BiTree)
			rc.Val = series[i+1]
			if rc.Val != nil {
				node.rchild = rc
				q.EnQueue(node.rchild)
			}
		}
	}
	return node
}

func (t *BiTree) LevelTraversal() []interface{} {
	series := make([]interface{}, 0)
	q := new(queue.Queue)
	q.InitQueue()
	if t != nil {
		q.EnQueue(t)
	}
	for ;!q.IsEmpty(); {
		node := q.DeQueue().(*BiTree)
		series = append(series, node.Val)
		if node.lchild!= nil {
			q.EnQueue(node.lchild)
		}
		if node.rchild != nil {
			q.EnQueue(node.rchild)
		}
	}
	return series
}

func PreOrderTraversal(root *BiTree) {
	if root == nil{
		return
	}
	Visit(root)
	PreOrderTraversal(root.lchild)
	PreOrderTraversal(root.rchild)
}

func InOrderTraversal(root *BiTree) []interface{} {
	if root != nil {
		InOrderTraversal(root.lchild)
		Visit(root)
		InOrderTraversal(root.rchild)
	}
	return nil
}

func PostOrderTraversal(root *BiTree)[]interface{} {
	if root != nil {
		PostOrderTraversal(root.lchild)
		PostOrderTraversal(root.rchild)
		Visit(root)
	}
	return nil
}

func Visit(node *BiTree) {
	fmt.Print(node.Val, " ")
}