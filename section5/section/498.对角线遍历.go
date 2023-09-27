package section

import "fmt"

/*
 *
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	isUpper, index, m, n := true, 0, len(mat), len(mat[0])
	if m == 0 || n == 0 {
		return nil
	}

	ans := make([]int, m*n)
	currentX, currentY := 0, 0

	next := func() {
		if isUpper {
			currentX -= 1
			currentY += 1
		} else {
			currentX += 1
			currentY -= 1
		}
	}

	for index < m*n {
		if currentX >= 0 && currentX < m && currentY >= 0 && currentY < n {
			ans[index] = mat[currentX][currentY]
			index += 1
			next()
		} else {
			if isUpper {
				currentX += 1
				isUpper = false

				if currentY >= n {
					next()
				}

			} else {
				currentY += 1
				isUpper = true
				if currentX >= m {
					next()
				}
			}
		}
	}

	return ans
}

// @lc code=end

func Test498() {
	// [1,2,4,7,5,3,6,8,9]
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	// [1,2,3,4]
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2},
		{3, 4},
	}))
}
