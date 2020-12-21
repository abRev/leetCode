package main

import (
	"fmt"
)

func decodeString(s string) string {
	times := 0
	start := 0
	bytes := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			timesArr = append(timesArr, times)
			times = 0
			bytes = append(bytes,s[i])
		} else if s[i] == ']' {
			sNew := []byte{}
		} 
	}

	return ""
}

func test() {
	if s[i] < 58 && s[i] > 47 {
			
		} else {
			bytes = append(bytes,s[i])
		}

	for times>0 {
		bytes = append(bytes,sNew...)
		times--;
	}
	bytes = []byte{}
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(s[1], s[len(s)-1])
	fmt.Println(decodeString(s))
}
