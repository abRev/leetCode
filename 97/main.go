package main

import (
	"fmt"
)

func main() {
	fmt.Println(isInterleave2("a", "", "c"))
}

// 空间改进版
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	second := []bool{}
	first := []bool{true}
	status := true
	for i, v := range s1 {
		if status && byte(v) == s3[i] {
			first = append(first, status)
		} else {
			status = false
			first = append(first, status)
		}
	}
	status = true
	for i, k := range s2 {
		second = []bool{}
		if status && s3[i] == byte(k) {
			second = append(second, status)
		} else {
			status = false
			second = append(second, status)
		}
		for i2, k2 := range s1 {
			if (s3[1+i+i2] == byte(k2) && second[i2]) || (first[i2+1] && s3[1+i+i2] == byte(k)) {
				second = append(second, true)
			} else {
				second = append(second, false)
			}
		}
		first = second
	}
	if len(second) > 0 {
		return second[len(second)-1]
	} else {
		if len(first) > 0 {
			return first[len(first)-1]
		}
		return true
	}
}

func isInterleave(s2 string, s1 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	arr := make([][]bool, len(s2)+1)
	row := []bool{true}
	status := true
	for i, v := range s1 {
		if status && byte(v) == s3[i] {
			row = append(row, status)
		} else {
			status = false
			row = append(row, status)
		}
	}
	arr[0] = row
	status = true
	for i, k := range s2 {
		row = []bool{}
		if status && s3[i] == byte(k) {
			row = append(row, status)
		} else {
			status = false
			row = append(row, status)
		}
		for i2, k2 := range s1 {
			if (s3[1+i+i2] == byte(k2) && row[i2]) || (arr[i][i2+1] && s3[1+i+i2] == byte(k)) {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		arr[i+1] = row
	}
	return arr[len(s2)][len(s1)]
	// 错误的累加算法
	// 	result := []uint8{}
	// I:
	// 	for x := 0; x < len(s3); {
	// 		if x > len(s3)-1 {
	// 			break
	// 		}
	// 		for y := 0; y < len(s1); {
	// 			if s3[x] == s1[y] {
	// 				x++
	// 				y++
	// 			} else {
	// 				result = append(result, s3[x])
	// 				x++
	// 			}
	// 		}
	// 		result = append(result, s3[x:]...)
	// 		break I
	// 	}
	// 	fmt.Println(string(result))
	// 	return s2 == string(result)
}
