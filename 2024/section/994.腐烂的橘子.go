package section

import "fmt"

/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	step := 0

	stack := make([][2]int, 0, len(grid)*len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				stack = append(stack, [2]int{i, j})
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	i, limit := 0, len(stack)
	for i < len(stack) {
		x, y := stack[i][0], stack[i][1]
		for _, direction := range directions {
			newX, newY := x+direction[0], y+direction[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 1 {
				grid[newX][newY] = 2
				stack = append(stack, [2]int{newX, newY})
			}
		}

		i += 1
		if i == limit && len(stack) > limit {
			step += 1
			limit = len(stack)
		}
	}

	for _, row := range grid {
		for _, column := range row {
			if column == 1 {
				return -1
			}
		}
	}

	return step
}

// @lc code=end

func Test994() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}) == 4)
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}) == -1)
	fmt.Println(orangesRotting([][]int{{0, 2}}) == 0)
}
