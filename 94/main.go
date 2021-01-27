package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return dfs(root)
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, dfs(root.Left)...)
	result = append(result, root.Val)
	result = append(result, dfs(root.Right)...)
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	result := inorderTraversal(root)
	fmt.Println(result)
}
