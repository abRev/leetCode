package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	length := 0
	for h := root; h != nil; h = h.Next {
		length++
	}
	remainder := length % k
	row := length / k
	result := []*ListNode{}
	for i := 0; i < k; i++ {
		start := root
		result = append(result, start)
		middle := row
		if remainder > 0 {
			middle++
			remainder--
		}
		for j := 0; j < middle-1; j++ {
			root = root.Next
		}
		if root != nil {
			end := root.Next
			root.Next = nil
			root = end
		}
	}
	return result
}

func main() {
	root := &ListNode{
		Val: 1,
	}
	root.Next = &ListNode{
		Val: 2,
	}
	fmt.Println(splitListToParts(root, 5))
}
