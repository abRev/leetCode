package main

import (
	"fmt"
)

func getPos(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func main() {
	nums := []int{
		2, 2, 3, 4, 6, 7, 7, 8, 9, 10,
	}
	fmt.Println(getPos(nums, 7))
	// hp := heap.NewWithIntComparator(1, 2, 3, 4)
	// fmt.Println(hp.Pop())
	// hp.Push(5)
	// fmt.Println(hp.Pop())
	// fmt.Println(hp.Pop())
	// fmt.Println(hp.Pop())
	// fmt.Println(hp.Pop())
}
