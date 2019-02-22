package structure

import "fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func createTree(series []int) *TreeNode {
	root := new(TreeNode)
	fmt.Printf("%T", root)
	//for i:=0;i<len(series);i+=2{
	//	layer := []TreeNode{}
	//
	//
	//}
	//node := new(TreeNode)
	return root
}

func createTreeByLevel(series []int) *TreeNode {
	queue := []int
}


func main() {
	series := []int{1, 2, 3, 4, 5}
	createTree(series)
}
