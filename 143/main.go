package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	arr := []*ListNode{}
	h := head
	for h != nil {
		arr = append(arr, h)
		h = h.Next
	}
	if len(arr) < 3 {
		return
	}
	middle := len(arr) / 2
	for i := len(arr) - 1; i > middle; i-- {
		arr[i].Next = arr[len(arr)-i-1].Next
		arr[len(arr)-i-1].Next = arr[i]
	}
	arr[middle].Next = nil
}

func main() {
	head := &ListNode{
		Val: 0,
	}
	head.Next = &ListNode{
		Val: 1,
	}
	head.Next.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}
	reorderList(head)
	fmt.Println("----")
}
