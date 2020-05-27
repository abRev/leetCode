package main

import "fmt"

// 暴力前缀和
func subarraysDivByK(A []int, K int) int {
	nums := []int{0}
	count := 0
	for _, val := range A {
		count += val
		nums = append(nums, count)
	}
	result := 0
	for i, val := range nums {
		for j := 0; j < i; j++ {
			if (val-nums[j])%K == 0 {
				result++
			}
		}
	}
	return result
}

// 如果两个数之差能被K整除，那么这两个数对k取余后之差一定为零

func subarraysDivByKByHash(A []int, K int) int {
	numMap := make(map[int]int)
	numMap[0] = 1
	count := 0
	result := 0
	for _, val := range A {
		count = (count + val) % K
		if count < 0 {
			count = count + K
		}
		if val, ok := numMap[count]; ok == true {
			result += val
		}
		if row, ok := numMap[count]; ok == true {
			numMap[count] = row + 1
		} else {
			numMap[count] = 1
		}
	}
	return result
}

func main() {
	nums := []int{
		4, 5, 0, -2, -3, 1,
	}
	fmt.Println(subarraysDivByKByHash(nums, 5))
}
