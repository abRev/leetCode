package main

import "fmt"

/*
从第一个公式到最后一个公式
依次加入到判别式（hashMap）中
每加入一个属性，则遍历全部公式，找到与属性相关联的属性，如果与已知不同，则false 相同则不处理，没有则加入
*/
func equationsPossible(equations []string) bool {
	h1 := make(map[byte]bool)
	h2 := make(map[byte]bool)
	for i1, row := range equations {
		v1, o1 := h1[row[0]]
		v2, o2 := h2[row[0]]
		v3, o3 := h1[row[3]]
		v4, o4 := h2[row[3]]
		if o1 == false || o2 == false {
			if o3 == false || o4 == false {
				// 当前公式两边都没有处理过
				if row[1] == '=' {
					h1[row[0]] = true
					h1[row[3]] = true
					for i2, row1 := range equations {
						if i1 == i2 {
							continue
						}
						if row1[0] == row[0] {
							if row1[1] == '=' {
								if _,ok := h2[row]
							} else {

							}
							
						}
					}
				} else {

				}
			}
		}
	}
	return true
}

func main() {
	equations := []string{
		"a==b", "b!=c", "c==a",
	}
	fmt.Println(equationsPossible(equations))
}
