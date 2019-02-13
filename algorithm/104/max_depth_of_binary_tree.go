package _04

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		return max(left, right) + 1
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
