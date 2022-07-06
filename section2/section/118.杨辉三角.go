package section

import "fmt"

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	rows := make([][]int, 0, numRows)
	rows = append(rows, []int{1})
	for length := 1; length < numRows; length++ {
		row := make([]int, 0, length+1)
		for i := 0; i <= length; i++ {
			column := 0
			if i-1 >= 0 {
				column += rows[length-1][i-1]
			}

			if i < length {
				column += rows[length-1][i]
			}
			row = append(row, column)
		}

		rows = append(rows, row)
	}

	return rows
}

func Test118() {
	// [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Println(generate(5))
	// [[1]]
	fmt.Println(generate(1))
}

// @lc code=end
