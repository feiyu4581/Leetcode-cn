package section

import "fmt"

/*
 * @lc app=leetcode.cn id=1992 lang=golang
 *
 * [1992] 找到所有的农场组
 */

// @lc code=start
func findFarmland(land [][]int) [][]int {
	res := make([][]int, 0)

	isFarm := func(i, j int) bool {
		if i < 0 || j < 0 {
			return false
		}

		if i >= len(land) || j >= len(land[i]) {
			return false
		}

		return land[i][j] == 1
	}

	isHead := func(i, j int) bool {
		if !isFarm(i, j) {
			return false
		}

		if i > 0 && land[i-1][j] == 1 {
			return false
		}

		if j > 0 && land[i][j-1] == 1 {
			return false
		}

		return true
	}

	findTail := func(i, j int) (int, int) {
		currentI, currentJ := i, j

		for isFarm(i, currentJ+1) {
			currentJ++
		}

		for isFarm(currentI+1, j) {
			currentI++
		}

		return currentI, currentJ
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if isHead(i, j) {
				tailI, tailJ := findTail(i, j)
				res = append(res, []int{i, j, tailI, tailJ})
			}
		}
	}

	return res
}

// @lc code=end

func Test1992() {
	// [[0,0,0,0],[1,1,2,2]]
	fmt.Println(findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	// [[0,0,1,1]]
	fmt.Println(findFarmland([][]int{{1, 1}, {1, 1}}))
	// []
	fmt.Println(findFarmland([][]int{{0}}))
	// [[1,0,1,1]]
	fmt.Println(findFarmland([][]int{{0, 0}, {1, 1}}))
	// [[0,1,1,1]]
	fmt.Println(findFarmland([][]int{{0, 1}, {0, 1}}))
}
