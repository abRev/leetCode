package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	return validPalindromeDeep(s, 0)
}

func validPalindromeDeep(s string, mark int) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			if mark > 0 {
				return false
			}
			mark++
			left := validPalindromeDeep(s[start+1:end+1], mark+1)
			right := validPalindromeDeep(s[start:end], mark+1)
			if left || right {
				return true
			} else {
				return false
			}
		}
		start++
		end--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
}
