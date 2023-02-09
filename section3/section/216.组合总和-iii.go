package section

import "fmt"

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtrackSum3 func(int, int, int, []int)
	backtrackSum3 = func(index, num, sum int, histories []int) {
		if num == 0 || index > 9 || index > sum {
			return
		}

		if num == 1 && sum >= index && sum <= 9 {
			copyHistories := make([]int, 0, k)
			for _, historyNum := range histories {
				copyHistories = append(copyHistories, historyNum)
			}

			copyHistories = append(copyHistories, sum)
			res = append(res, copyHistories)
			return
		}

		backtrackSum3(index+1, num, sum, histories)

		histories = append(histories, index)
		backtrackSum3(index+1, num-1, sum-index, histories)
		histories = histories[:len(histories)-1]
	}

	backtrackSum3(1, k, n, make([]int, 0, k))
	return res
}

func Test216() {
	// [[1,2,4]]
	fmt.Println(combinationSum3(3, 7))
	// [[1,2,6], [1,3,5], [2,3,4]]
	fmt.Println(combinationSum3(3, 9))
	// []
	fmt.Println(combinationSum3(4, 1))
}

// @lc code=end
