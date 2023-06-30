package section

import "fmt"

/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */

// @lc code=start
func countBattleships(board [][]byte) int {
	res := 0
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				res++
			}
		}
	}

	return res
}

// @lc code=end

func Test419() {
	fmt.Println(countBattleships([][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}) == 2)
	fmt.Println(countBattleships([][]byte{
		{'.'},
	}) == 0)
}
