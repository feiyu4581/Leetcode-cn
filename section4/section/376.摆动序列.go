package section

import "fmt"

/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	dp := make([][2]int, len(nums))
	for i := range dp {
		dp[i] = [2]int{1, 1}
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}

		return j
	}

	maxLength := dp[0][0]
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
				if dp[i][0] > maxLength {
					maxLength = dp[i][0]
				}
			} else if nums[i] < nums[j] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
				if dp[i][1] > maxLength {
					maxLength = dp[i][1]
				}
			}
		}
	}

	return maxLength
}

// @lc code=end

func Test376() {
	fmt.Println(wiggleMaxLength([]int{0, 0}) == 1)
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}) == 6)
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}) == 7)
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) == 2)
}
