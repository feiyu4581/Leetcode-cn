package section

import "fmt"

/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}
	if stones[1]-stones[0] > 1 {
		return false
	}

	cache := make(map[[2]int]bool, 0)

	var dfs func(int, int) bool
	dfs = func(i, step int) bool {
		if step == 0 {
			return false
		}

		if i == len(stones)-1 {
			return true
		}

		if res, ok := cache[[2]int{i, step}]; ok {
			return res
		}

		for j := i + 1; j < len(stones); j++ {
			diff := stones[j] - stones[i]
			if diff > step+1 {
				break
			}

			if diff >= step-1 && diff <= step+1 && dfs(j, diff) {
				cache[[2]int{i, step}] = true
				return true
			}
		}

		cache[[2]int{i, step}] = false
		return false
	}

	return dfs(1, 1)
}

// @lc code=end

func Test403() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}) == false)
}
