package section

import "fmt"

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	current := 0
	for _, num := range nums {
		if num == 1 {
			current += 1
			if current > ans {
				ans = current
			}
		} else {
			current = 0
		}
	}

	return ans
}

// @lc code=end

func Test485() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}) == 3)
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}) == 2)
}
