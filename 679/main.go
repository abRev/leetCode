package main

import (
	"fmt"
)

func judgePoint24(nums []int) bool {
	result := false
	newNums := []float64{}
	for k := 0; k < len(nums); k++ {
		newNums = append(newNums, float64(nums[k]))
	}
	deepFunc(newNums, &result)
	return result
}

func deepFunc(nums []float64, result *bool) {
	fmt.Println(nums)
	if (*result) == true {
		return
	}
	if len(nums) == 1 {
		if nums[0] == 24 {
			(*result) = true
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			// 生成一个新数组
			newNums := []float64{}
			for k := 0; k < len(nums); k++ {
				if i == k || j == k {
					continue
				}
				newNums = append(newNums, nums[k])
			}
			// +
			value := nums[i] + nums[j]
			deepFunc(append(newNums, value), result)
			if (*result) == true {
				return
			}
			// -
			value = nums[i] - nums[j]
			deepFunc(append(newNums, value), result)
			if (*result) == true {
				return
			}
			// *
			value = nums[i] * nums[j]
			deepFunc(append(newNums, value), result)
			if (*result) == true {
				return
			}
			// /
			if nums[j] != 0 {
				value = nums[i] / nums[j]
				deepFunc(append(newNums, value), result)
			}
			if (*result) == true {
				return
			}
		}
	}
}

func main() {
	nums := []int{3, 8, 3, 8}
	fmt.Println(judgePoint24(nums))
}
