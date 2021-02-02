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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	index := indexOf(inorder, postorder[len(postorder)-1])
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func indexOf(arr []int, key int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

func main() {
	left := []int{9, 3, 15, 20, 7}
	right := []int{9, 15, 7, 20, 3}
	result := buildTree(left, right)
	fmt.Println(result)
}
