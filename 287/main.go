package main

import "fmt"

func findDuplicate(nums []int) int {
	right := len(nums)
	left := 0
	length := 0
	for left < right {
		length = 0
		mid := (left + right) / 2
		for _, val := range nums {
			if val <= mid {
				length++
			}
		}
		if length > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{
		3, 1, 3, 4, 2,
	}
	fmt.Println(findDuplicate(nums))
}
