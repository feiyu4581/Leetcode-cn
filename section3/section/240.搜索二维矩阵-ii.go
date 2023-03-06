package section

import "fmt"

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start

func searchMatrixByPosition(matrix [][]int, startX, startY, endX, endY, target int) bool {
	if startX > endX || startY > endY {
		return false
	}

	if target < matrix[startX][startY] || target > matrix[endX][endY] {
		return false
	}

	if startX == endX && startY == endY {
		return matrix[startX][startY] == target
	}

	midX, midY := startX+(endX-startX)>>1, startY+(endY-startY)>>1
	if matrix[midX][midY] == target {
		return true
	} else if matrix[midX][midY] > target {
		return searchMatrixByPosition(matrix, startY, startY, midX-1, endY, target) ||
			searchMatrixByPosition(matrix, midX, startY, endX, midY-1, target)
	} else {
		return searchMatrixByPosition(matrix, startX, midY+1, midX, endY, target) ||
			searchMatrixByPosition(matrix, midX+1, startY, endX, endY, target)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	return searchMatrixByPosition(matrix, 0, 0, len(matrix)-1, len(matrix[0])-1, target)
}

// @lc code=end

func Test240() {
	fmt.Println(searchMatrix([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}, 5))

	fmt.Println(searchMatrix([][]int{
		{1, 4},
		{2, 5},
	}, 3) == false)

	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5))

	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20) == false)
}
