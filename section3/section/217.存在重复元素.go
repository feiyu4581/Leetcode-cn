package section

import "fmt"

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}

func Test217() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}) == false)
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

// @lc code=end
