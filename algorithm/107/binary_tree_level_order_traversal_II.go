package _07

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var levelList [][]int
	if root == nil {
		return nil
	}
	//先将根节点添加至List
	var fatherList []*TreeNode
	fatherList = append(fatherList, root)
	for {
		if len(fatherList) == 0 {
			break
		}
		var sonList []*TreeNode
		var sonValList []int
		for i := 0; i < len(fatherList); i++ {
			sonValList = append(sonValList, fatherList[i].Val)
			if fatherList[i].Left != nil {
				sonList = append(sonList, fatherList[i].Left)
			}
			if fatherList[i].Right != nil {
				sonList = append(sonList, fatherList[i].Right)
			}
		}
		fatherList = sonList
		//tmpList := levelList[:]
		//levelList = append([][]int{}, sonValList)
		//levelList = append(levelList,  tmpList...)
		levelList = append(levelList, sonValList)
	}
	return levelList

}
