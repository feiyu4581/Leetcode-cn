package section

import "fmt"

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start

func rangeNumIslands(i, j int, grid [][]byte, visited [][]bool) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' || visited[i][j] {
		return false
	}

	visited[i][j] = true
	res := true

	res = rangeNumIslands(i+1, j, grid, visited) || res
	res = rangeNumIslands(i-1, j, grid, visited) || res
	res = rangeNumIslands(i, j+1, grid, visited) || res
	res = rangeNumIslands(i, j-1, grid, visited) || res

	return res
}

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if rangeNumIslands(i, j, grid, visited) {
				count += 1
			}
		}
	}

	return count
}

func Test200() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}) == 1)

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}) == 3)

	fmt.Println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}) == 1)
}

// @lc code=end
