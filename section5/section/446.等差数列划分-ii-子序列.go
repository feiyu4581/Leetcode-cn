package section

import "fmt"

/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	f := make([]map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = make(map[int]int)
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - nums[i]
			ans += f[i][diff]
			f[j][diff] += f[i][diff] + 1
		}
	}

	return ans
}

// @lc code=end

func Test446() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}) == 7)
	fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}) == 16)
}
