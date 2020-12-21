package main

/**
计算两个字符串的距离
动态规划
两个字符有四种情况(a,b)
a b 都是字符
a 是字符，b是空格
a是空格，b是字符
a b 都是空格
第四种情况排除掉
第二第三种情况距离是左，上，（上左） 最小值+1
第一种情况分相等不相等，相等则等于左上，不等则同第二种情况
*/

import (
	"fmt"
)

func lci(str1, str2 string) {
	ho := make([][]int, len(str2)+1)
	for i := 0; i <= len(str2); i++ {
		row := make([]int, len(str1)+1)
		ho[i] = row
	}
	for i := 0; i < len(str1)+1; i++ {
		ho[0][i] = i
	}
	for i := 0; i < len(str2)+1; i++ {
		ho[i][0] = i
	}
	for i := 1; i <= len(str2); i++ {
		for j := 1; j <= len(str1); j++ {
			if str2[i-1] == str1[j-1] {
				ho[i][j] = ho[i-1][j-1]
			} else {
				ho[i][j] = min(ho[i-1][j], ho[i][j-1], ho[i-1][j-1]) + 1
			}
		}
	}
	for _, row := range ho {
		fmt.Println(row)
	}
}

func min(n1, n2, n3 int) int {
	if n1 > n2 {
		if n2 > n3 {
			return n3
		} else {
			return n2
		}
	} else {
		if n1 > n3 {
			return n3
		} else {
			return n1
		}
	}
}

func main() {
	str1 := "abangma"
	str2 := "cabaga"
	lci(str1, str2)
}
