package tree

import (
	"fmt"
	"testing"
)

func TestCreateTreeByLevel(t *testing.T) {
	tree := new(BiTree)
	series := []interface{}{1,23,4,nil,5,6,12}
	tree = CreateTreeByLevel(series)
	fmt.Println(tree.LevelTraversal())
	PreOrderTraversal(tree)
	fmt.Println("\n----------------------------")
	InOrderTraversal(tree)
	fmt.Println("\n----------------------------")
	PostOrderTraversal(tree)
	fmt.Println("\n----------------------------")
}