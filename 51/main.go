package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	left := newInterval[0]
	right := newInterval[1]
	start := 0
	end := len(intervals) - 1
	leftNeed := -1
	rightNeed := len(intervals)
	for key, value := range intervals {
		if value[0] > left {
			leftNeed = key - 1
			break
		}
		if value[0] <= left && left <= value[1] {
			leftNeed = key - 1
			break
		}
		leftNeed = key
	}
	for key, value := range intervals {
		if value[0] > right {
			rightNeed = key
			break
		}
		if value[0] <= right && right <= value[1] {
			rightNeed = key + 1
			break
		}
		rightNeed = key + 1
	}
	result := [][]int{}
	// 组装前半段
	for i := 0; i <= leftNeed; i++ {
		result = append(result, intervals[i])
	}
	// 插入中间一段
	if leftNeed+1 < 0 || leftNeed+1 >= len(intervals) {
		start = left
	} else if left < intervals[leftNeed+1][0] {
		start = left
	} else if intervals[leftNeed+1][0] <= left && left <= intervals[leftNeed+1][1] {
		start = intervals[leftNeed+1][0]
	}
	if rightNeed-1 < 0 || rightNeed-1 >= len(intervals) {
		end = right
	} else if right > intervals[rightNeed-1][1] {
		end = right
	} else {
		end = intervals[rightNeed-1][1]
	}

	result = append(result, []int{start, end})
	// 组装后半段
	for i := rightNeed; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}
	return result
}

func main1() {
	intervals := [][]int{{0, 5}, {9, 12}}
	newInterval := []int{7, 16}
	fmt.Println(insert(intervals, newInterval))
}
