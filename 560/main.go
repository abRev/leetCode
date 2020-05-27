package main

import "fmt"

func subarraySum(nums []int, k int) int {
	arr := []int{0}
	count := 0

	for _, val := range nums {
		count += val
		arr = append(arr, count)
	}
	count = 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == k {
				count++
			}
		}
	}
	return count
}

func subarraySumByHash(nums []int, k int) int {
	numMaps := make(map[int]int)
	numMaps[0] = 1
	count := 0
	result := 0
	for _, val := range nums {
		count += val
		if _, ok := numMaps[count-k]; ok == true {
			result += numMaps[count-k]
		}
		if v, ok := numMaps[count]; ok == true {
			numMaps[count] = v + 1
		} else {
			numMaps[count] = 1
		}
	}
	return result
}

func main() {
	nums := []int{1}
	fmt.Println(subarraySumByHash(nums, 0))
}
