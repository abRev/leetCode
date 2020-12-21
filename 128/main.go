package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	mapArr := make(map[int]bool)
	for _, row := range nums {
		if _, ok := mapArr[row]; ok == false {
			mapArr[row] = true
		}
	}
	result := 1
	for key, _ := range mapArr {
		y := 1
		for {
			if _, ok := mapArr[key+y]; ok == true {
				y++
			} else {
				break
			}
		}
		if result < y {
			result = y
		}
		y = 1
	}
	return result
}

func main() {
	nums := []int{0, -1}
	fmt.Println(longestConsecutive(nums))
}
