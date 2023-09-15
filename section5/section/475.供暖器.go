package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */

// @lc code=start

func min475(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max475(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	res := 0
	current := 0
	for i := 0; i < len(houses); i++ {
		for current < len(heaters)-1 && houses[i] > heaters[current] {
			current += 1
		}

		if houses[i] >= heaters[current] {
			res = max475(res, houses[i]-heaters[current])
		} else {
			right := heaters[current] - houses[i]
			if current > 0 {
				left := houses[i] - heaters[current-1]
				res = max475(res, min475(left, right))
			} else {
				res = max475(res, right)
			}
		}
	}

	return res
}

// @lc code=end

func Test475() {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}) == 1)
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}) == 1)
	fmt.Println(findRadius([]int{1, 5}, []int{2}) == 3)
}
