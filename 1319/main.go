package main

import "fmt"

// 深度优先
func MakeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	arr := make([][]int, n)
	for i := 0; i < len(connections); i++ {
		x, y := connections[i][0], connections[i][1]
		arr[x] = append(arr[x], y)
		arr[y] = append(arr[y], x)
	}
	result := map[int]bool{}
	var dfs func(int)
	dfs = func(n int) {
		result[n] = true
		for _, val := range arr[n] {
			if !result[val] {
				dfs(val)
			}
		}
	}
	size := 0
	for i := 0; i < n; i++ {
		if !result[i] {
			size++
			dfs(i)
		}
	}
	return size - 1
}

// 广度优先
func MakeConnected2(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	arr := make([][]int, n)
	for _, val := range connections {
		x, y := val[0], val[1]
		arr[x] = append(arr[x], y)
		arr[y] = append(arr[y], x)
	}
	result := 0
	ais := map[int]bool{}
	for i := 0; i < n; i++ {
		if ais[i] == false {
			result++
			ais[0] = true
		} else {
			continue
		}
		row := arr[i]
		for len(row) > 0 {
			temp := []int{}
			for _, val := range row {
				if ais[val] == false {
					ais[val] = true
					for _, val1 := range arr[val] {
						temp = append(temp, val1)
					}
				}
			}
			row = temp
		}
	}
	return result - 1
}

func main() {
	result := MakeConnected2(6, [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 3}, []int{1, 2}})
	fmt.Println(result)
}
