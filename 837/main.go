package main

import "fmt"

func new21Game(N int, K int, W int) float64 {
	// 实际计算的是[K，N] (N，K-1+W] 的比例的，只需要求出组合出每个数字用到的可能 求和求占比就行了
	// 计算每个数有第三种组合，通过动态规划法，W* k-1+W 的数组，每一层求数量
	// 然后分别求出 第 W行 的所有数求和求比例
	nums := [][]int{}
	num1 := []int{}
	for i := 0; i <= K-1+W; i++ {
		num1 = append(num1, 1)
	}
	nums = append(nums, num1)
	for i := 1; i < W; i++ {
		num := []int{1}
		for j := 1; j <= K-1+W; j++ {
			if i > j {
				num = append(num, nums[i-1][j])
			} else {
				num = append(num, nums[i-1][j]+num[j-i])
			}
		}
		nums = append(nums, num)
	}
	low := 0
	count := 0
	for i := K; i <= K-1+W; i++ {
		if i <= N {
			low += nums[W-1][i]
		}
		count += nums[W-1][i]
	}
	return float64(low) / float64(count)
}

func main() {
	fmt.Println(new21Game(21, 17, 10))
}
