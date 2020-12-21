package main

import "fmt"

func addBinary(a string, b string) string {
	aArr := []byte(a)
	bArr := []byte(b)
	sizeA := len(a)
	sizeB := len(b)
	var sizeMin, sizeMax int
	if sizeA > sizeB {
		sizeMin = sizeB
		sizeMax = sizeA
	} else {
		sizeMin = sizeA
		sizeMax = sizeB
	}
	result := make([]byte, sizeMax)
	mark := 0
	for i := sizeMin - 1; i >= 0; i-- {
		if aArr[i] == '1' {
			if bArr[i] == '1' {
				if mark == 1 {
					result[i] = '1'
					mark = 1
				} else {
					result[i] = '0'
					mark = 1
				}
			} else {
				if mark == 1 {
					result[i] = '0'
					mark = 1
				} else {
					result[i] = '1'
					mark = 0
				}
			}
		} else {
			if bArr[i] == '1' {
				if mark == 1 {
					result[i] = '0'
					mark = 1
				} else {
					result[i] = '1'
					mark = 0
				}
			} else {
				if mark == 1 {
					result[i] = '1'
					mark = 0
				} else {
					result[i] = '0'
					mark = 0
				}
			}
		}
	}
	if mark == 1 {
		newResult := []byte{'1'}
		result = append(newResult, result...)
	}
	return string(result[:])
}

func main() {
	a := "1100"
	b := "100"
	fmt.Println(addBinary(a, b))
}
