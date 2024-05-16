package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=1953 lang=golang
 *
 * [1953] 你可以工作的最大周数
 */

// @lc code=start
func numberOfWeeks(milestones []int) int64 {
	allNums, max := 0, 0
	for _, num := range milestones {
		allNums += num
		if num > max {
			max = num
		}
	}

	if otherNums := allNums - max; otherNums < max {
		return int64(otherNums*2 + 1)
	} else {
		return int64(allNums)
	}
}

// @lc code=end

func Test1953() {
	fmt.Println(numberOfWeeks([]int{1, 2, 3}) == 6)
	fmt.Println(numberOfWeeks([]int{5, 2, 1}) == 7)
	fmt.Println(numberOfWeeks([]int{5, 7, 5, 7, 9, 7}) == 40)
}
