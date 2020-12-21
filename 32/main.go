package main

import (
	"fmt"
	"leetcode/stack"
)

func longestValidParentheses(s string) int {
	stack := stack.New('0')
	result := 0
	row :=0
	for _,val := range s {
		if val == ')' {
			head:= stack.Head()
			if head == '(' {
				stack.Pop()
				row += 2
			} else {
				if row > result {
					result = row
				}
				row = 0
				stack.Push(val)
			}
		} else {
			stack.Push(val)
		}
	}
	if row > result {
		result = row
	}
	return result
}

func main(){
	fmt.Println(longestValidParentheses("()(()"))
}