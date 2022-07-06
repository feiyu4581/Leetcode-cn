package section

import "fmt"

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start

func computeNextRow(row []int, index int) []int {
	nextRow := make([]int, 0, len(row)+1)
	for index := 0; index <= len(row); index++ {
		column := 0
		if index-1 >= 0 {
			column += row[index-1]
		}

		if index < len(row) {
			column += row[index]
		}

		nextRow = append(nextRow, column)
	}

	if index == 0 {
		return nextRow
	}

	return computeNextRow(nextRow, index-1)
}

func getRow(rowIndex int) []int {
	startRow := []int{1}
	if rowIndex == 0 {
		return startRow
	}

	return computeNextRow(startRow, rowIndex-1)
}

func Test119() {
	// [1,3,3,1]
	fmt.Println(getRow(3))

	// [1]
	fmt.Println(getRow(0))

	// [1, 1]
	fmt.Println(getRow(1))
}

// @lc code=end
