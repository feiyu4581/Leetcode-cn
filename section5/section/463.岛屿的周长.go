package section

import "fmt"

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x < 0 || x >= len(grid) || y < 0 || y >= len(row) || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return ans
}

// @lc code=end

func Test463() {
	fmt.Println(islandPerimeter([][]int{{1}}) == 4)
	fmt.Println(islandPerimeter([][]int{{1, 0}}) == 4)
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}) == 16)
}
