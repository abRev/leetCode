package main

import "fmt"

func numIslands(grid [][]byte) int {
	// for i:=0;i<len(grid[0]);i++ {
	//     if grid[0][i] == '1' {
	//         deleteHalfIslands(grid, 0,i)
	//     }
	//     if grid[len(grid)-1][i] == '1' {
	//         deleteHalfIslands(grid, len(grid)-1,i)
	//     }
	// }
	// for i:=0; i< len(grid);i++ {
	//     if grid[i][0] == '1' {
	//         deleteHalfIslands(grid, i,0)
	//     }
	//     if grid[i][len(grid)-1] == '1' {
	//         deleteHalfIslands(grid, i,len(grid)-1)
	//     }
	// }
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				deleteHalfIslands(grid, i, j)
			}
		}
	}
	return result
}

// 岛屿删除
func deleteHalfIslands(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) {
		return
	}
	if y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	deleteHalfIslands(grid, x, y-1)
	deleteHalfIslands(grid, x, y+1)
	deleteHalfIslands(grid, x-1, y)
	deleteHalfIslands(grid, x+1, y)
}

func main() {
	// grid := [][]byte{
	// 	{'1','1','1','1','0'},
	// 	{'1','1','0','1','0'},
	// 	{'1','1','0','0','0'},
	// 	{'0','0','0','0','0'},
	// }
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
