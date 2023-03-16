package section

import "fmt"

/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

//如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
//如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
//如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
//如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

// @lc code=start

var positions = [][2]int{
	{1, 0},
	{1, -1},
	{1, 1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{0, -1},
	{0, 1},
}

func gameOfLife(board [][]int) {
	lifeReporting := func(i, j int) int {
		count := 0
		for _, position := range positions {
			x, y := i+position[0], j+position[1]
			if x >= 0 && y >= 0 && x < len(board) && y < len(board[0]) {
				if board[x][y]%2 == 1 {
					count += 1
				}
			}
		}

		return count
	}

	for i := range board {
		for j := range board[i] {
			life := lifeReporting(i, j)
			if board[i][j] == 1 {
				if life < 2 || life > 3 {
					board[i][j] = 3
				}
			} else {
				if life == 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

// @lc code=end

func Test289() {
	life := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(life)
	// [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
	fmt.Println(life)

	life = [][]int{
		{1, 1},
		{1, 0},
	}
	gameOfLife(life)
	// [[1,1],[1,1]]
	fmt.Println(life)
}
