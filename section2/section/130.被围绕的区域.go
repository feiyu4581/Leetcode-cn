package section

import "fmt"

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start

func recursiveSolve(i, j int, allowO [][]bool, board [][]byte, visisted [][]bool) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	if board[i][j] == 'X' {
		return
	}

	if visisted[i][j] {
		return
	}

	allowO[i][j] = true
	visisted[i][j] = true
	recursiveSolve(i-1, j, allowO, board, visisted)
	recursiveSolve(i+1, j, allowO, board, visisted)
	recursiveSolve(i, j+1, allowO, board, visisted)
	recursiveSolve(i, j-1, allowO, board, visisted)
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	allowO := make([][]bool, len(board))
	visisted := make([][]bool, len(board))
	for index := range allowO {
		allowO[index] = make([]bool, len(board[0]))
		visisted[index] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		allowO[i][0] = board[i][0] == 'O'
		allowO[i][len(board[0])-1] = board[i][len(board[0])-1] == 'O'
	}

	for j := 0; j < len(board[0]); j++ {
		allowO[0][j] = board[0][j] == 'O'
		allowO[len(board)-1][j] = board[len(board)-1][j] == 'O'
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if allowO[i][j] {
				recursiveSolve(i, j, allowO, board, visisted)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && !allowO[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func Test130() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(board)
	// {'X', 'X', 'X', 'X'},
	// {'X', 'X', 'X', 'X'},
	// {'X', 'X', 'X', 'X'},
	// {'X', 'O', 'X', 'X'},
	fmt.Println(board)
}

// @lc code=end
