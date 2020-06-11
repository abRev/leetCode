package main

import "fmt"

// 暴力法（开始以为可以用前缀和来做，写着写着发现使用前缀和跟从原始数据上操作没区别，还额外浪费时间）
func dailyTemperatures(T []int) []int {
	arr := []int{0}
	for i := 0; i < len(T)-1; i++ {
		arr = append(arr, T[i+1]-T[i])
	}
	arrNew := make([]int, len(T))
	count := 0
	for i := 0; i < len(arr); i++ {
		count += arr[i]
		arrNew[i] = count
	}
	result := make([]int, len(T))
	for i := 0; i < len(arrNew)-1; i++ {
		size := 1
		for j := i + 1; j < len(arrNew); j++ {
			if arrNew[j] > arrNew[i] {
				result[i] = size
				size = 1
				break
			} else {
				result[i] = 0
				size++
			}
		}
	}
	return result
}

// 使用栈来做（懒得写栈了，直接用数组代替，数组长度最高为数组长度）
func dailyTemperatures2(T []int) []int {
	stack := []int{}
	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

// 用切片怀疑有性能损耗，使用数组加i来实现看看
func dailyTemperatures3(T []int) []int {
	stack := make([]int, len(T))
	top := 0
	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for top > 0 && T[i] > T[stack[top-1]] {
			result[stack[top-1]] = i - stack[top-1]
			top--
		}
		stack[top] = i
		top++
	}
	return result
}

func main() {
	t := []int{
		55, 38, 53, 81, 61, 93, 97, 32, 43, 78,
	}
	fmt.Println(dailyTemperatures3(t))
}
