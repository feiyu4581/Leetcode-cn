package section

import "fmt"

/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	matrix [][]int
	cache  map[[4]int]int
}

func Constructor304(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
		cache:  make(map[[4]int]int, 0),
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	key := [4]int{row1, col1, row2, col2}
	if _, ok := this.cache[key]; !ok {
		res := 0
		for i := row1; i <= row2; i++ {
			for j := col1; j <= col2; j++ {
				res += this.matrix[i][j]
			}
		}

		this.cache[key] = res
	}

	return this.cache[key]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

func Test304() {
	numMatrix := Constructor304([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3) == 8)  // return 8 (红色矩形框的元素总和)
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2) == 11) // return 11 (绿色矩形框的元素总和)
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4) == 12) // return 12 (蓝色矩形框的元素总和)
}
