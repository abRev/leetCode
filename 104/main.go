package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	result := 0
	dfs(root, 0, &result)
	return result
}

func dfs(root *TreeNode, now int, deep *int) {
	if root == nil {
		return
	}
	now++
	if now > *deep {
		*deep = now
	}
	dfs(root.Left, now, deep)
	dfs(root.Right, now, deep)
}
