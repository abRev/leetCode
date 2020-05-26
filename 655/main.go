package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	length := 0
	getLength(root, &length, 1)
	arr := make([][]string, length)
	size := int(math.Pow(float64(2), float64(length)) - 1)
	for i := 0; i < length; i++ {
		arr[i] = make([]string, size)
	}
	deepSetArr(&arr, (size-1)/2, 0, (size-1)/2, root)
	return arr
}

func deepSetArr(arr *[][]string, size, n int, index int, root *TreeNode) {
	size = (size - 1) / 2
	(*arr)[n][index] = strconv.Itoa(root.Val)
	if root.Left != nil {
		deepSetArr(arr, size, n+1, index-size-1, root.Left)
	}
	if root.Right != nil {
		deepSetArr(arr, size, n+1, index+1+size, root.Right)
	}
}

func getLength(root *TreeNode, length *int, now int) {
	if root.Left != nil {
		getLength(root.Left, length, now+1)
	}
	if root.Right != nil {
		getLength(root.Right, length, now+1)
	}
	if *length < now {
		*length = now
	}
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
	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}
	root.Right.Right = &TreeNode{
		Val: 6,
	}
	root.Left.Left.Left = &TreeNode{
		Val: 7,
	}
	fmt.Println(printTree(root))
}
