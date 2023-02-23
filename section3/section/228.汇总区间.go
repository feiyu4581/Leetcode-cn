package section

import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}

	start, lastNum := nums[0], nums[0]

	collectNum := func() {
		if start == lastNum {
			res = append(res, fmt.Sprintf("%d", start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, lastNum))
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNum+1 {
			lastNum = nums[i]
		} else {
			collectNum()
			start, lastNum = nums[i], nums[i]
		}
	}

	collectNum()
	return res
}

// @lc code=end

func Test228() {
	// ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	// ["0","2->4","6","8->9"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
