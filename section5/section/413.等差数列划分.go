package section

import "fmt"

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	diffs := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		diffs[i] = nums[i+1] - nums[i]
	}

	computeNums := func(num int) int {
		res := 0
		for i := 1; i < num; i++ {
			res += i
		}

		return res
	}

	continueNums := 1
	res := 0
	for i := 1; i < len(diffs); i++ {
		if diffs[i] == diffs[i-1] {
			continueNums += 1
		} else {
			res += computeNums(continueNums)
			continueNums = 1
		}
	}

	return res + computeNums(continueNums)
}

// @lc code=end

func Test413() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}) == 3)
	fmt.Println(numberOfArithmeticSlices([]int{1}) == 0)
}
