package main

import (
	"fmt"
)

func LocNum(str1, str2 string) ([][]int, int) {
	result := 0
	ho := make([][]int, len(str2)+1)
	for i := 0; i < len(str2)+1; i++ {
		row := make([]int, len(str1)+1)
		ho[i] = row
	}
	for i := 1; i <= len(str2); i++ {
		for j := 1; j < len(str1); j++ {
			if str1[j-1] == str2[i-1] {
				ho[i][j] = max(ho[i-1][j], ho[i][j-1], ho[i-1][j-1]) + 1
				if ho[i][j] > result {
					result = ho[i][j]
				}
			} else {
				ho[i][j] = 0
			}
		}
	}
	return ho, result
}

func max(n1, n2, n3 int) int {
	if n1 > n2 {
		if n1 < n3 {
			return n3
		} else {
			return n1
		}
	}
	if n2 > n3 {
		return n2
	}
	return n3
}

func main() {
	str1 := "abangma"
	str2 := "cabaga"
	result, max := LocNum(str1, str2)
	fmt.Print(" 0 0 ")
	for _, c := range str1 {
		fmt.Print(string(c), " ")
	}
	fmt.Println()
	for i, row := range result {
		if i == 0 {
			fmt.Print("0")
		} else {
			fmt.Print(string(str2[i-1]))

		}
		fmt.Println(row)
	}
	fmt.Println("max: ", max)
}
