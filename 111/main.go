package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	deepMin := -1
	deepNowInt := 0
	return deepNow(root, deepNowInt, deepMin)
}

func deepNow(root *TreeNode, deepNowInt int, deepMin int) int {
	if root == nil {
		return deepMin
	}
	deepNowInt++
	if root.Left == nil && root.Right == nil {
		if deepMin == -1 {
			deepMin = deepNowInt
		} else if deepMin > deepNowInt {
			deepMin = deepNowInt
		}
	}
	deepMin = deepNow(root.Left, deepNowInt, deepMin)
	deepMin = deepNow(root.Right, deepNowInt, deepMin)
	return deepMin
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 1,
	}
	root.Right.Right = &TreeNode{
		Val: 1,
	}
	root.Right.Right.Right = &TreeNode{
		Val: 1,
	}

	fmt.Println(minDepth(root))
}
