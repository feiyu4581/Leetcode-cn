package section

import "fmt"

/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start
func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	res := make(map[[2]int]struct{}, 0)

	min := func(i, j int) int {
		if i < j {
			return i
		}

		return j
	}

	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == targetCapacity || j == targetCapacity || i+j == targetCapacity {
			return true
		}

		if _, ok := res[[2]int{i, j}]; ok {
			return false
		}

		res[[2]int{i, j}] = struct{}{}

		if i != 0 && dfs(0, j) {
			return true
		}

		if j != 0 && dfs(i, 0) {
			return true
		}

		if i != jug1Capacity && dfs(jug1Capacity, j) {
			return true
		}

		if j != jug2Capacity && dfs(i, jug2Capacity) {
			return true
		}

		if i > 0 && j < jug2Capacity {
			if diffI := min(jug2Capacity-j, i); dfs(i-diffI, j+diffI) {
				return true
			}
		}

		if j > 0 && i < jug1Capacity {
			if diffJ := min(jug1Capacity-i, j); dfs(i+diffJ, j-diffJ) {
				return true
			}
		}

		return false
	}

	return dfs(0, 0)
}

// @lc code=end

func Test365() {
	fmt.Println(canMeasureWater(3, 5, 4))
	fmt.Println(canMeasureWater(2, 6, 5) == false)
	fmt.Println(canMeasureWater(1, 2, 3) == true)
}
