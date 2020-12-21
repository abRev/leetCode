package main

import "fmt"

/**
找到根节点。入度为0的节点为根节点，如果根节点的数量超过1个 则无效
判断所有点的入度，有一个点的入度超过1，则无效
*/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	inMap := make(map[int]int)
	for i := 0; i < n; i++ {
		inMap[i] = 0
	}
	for i := 0; i < n; i++ {
		if leftChild[i] > -1 {
			if value, ok := inMap[leftChild[i]]; ok == true {
				if value >= 1 {
					return false
				}
				inMap[leftChild[i]] = value + 1
			} else {
				inMap[leftChild[i]] = 1
			}
		}
		if rightChild[i] > -1 {
			if value, ok := inMap[rightChild[i]]; ok == true {
				if value >= 1 {
					return false
				}
				inMap[rightChild[i]] = value + 1
			} else {
				inMap[rightChild[i]] = 1
			}
		}
	}
	if len(inMap) < n {
		return false
	}
	inSize := 0
	for _, value := range inMap {
		if value == 0 {
			inSize++
		}
	}
	if inSize > 1 || inSize == 0 {
		return false
	}
	return true
}

func main() {
	left := []int{
		1, 2, 0, -1,
	}
	right := []int{
		-1, -1, -1, -1,
	}
	fmt.Println(validateBinaryTreeNodes(4, left, right))
}
