package main

import (
	"fmt"
)

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	result := []uint8{}
I:
	for x := 0; x < len(s3); {
		if x > len(s3)-1 {
			break
		}
		for y := 0; y < len(s1); {
			if s3[x] == s1[y] {
				x++
				y++
			} else {
				result = append(result, s3[x])
				x++
			}
		}
		result = append(result, s3[x:]...)
		break I
	}
	fmt.Println(string(result))
	return s2 == string(result)
}
