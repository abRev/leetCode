package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	for top < bottom {
		if matrix[bottom][0] > target {
			bottom--
		}
		if matrix[top][right] < target {
			top++
		}
	}
	if x == 0 {
		return false
	}
	for x > 0 {
		x--
		left := 0
		right := len(matrix[0]) - 1
		for left <= right {
			middle := (left + right) / 2
			if matrix[x][middle] == target {
				return true
			} else if matrix[x][middle] > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 20))
}
