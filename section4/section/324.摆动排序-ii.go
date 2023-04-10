package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	sort.Ints(nums)

	res := make([]int, len(nums))

	middle := (len(nums) + 1) / 2
	start := len(nums) - 1
	index := middle - 1
	i := 0
	for index >= 0 {
		res[i] = nums[index]

		if i+1 < len(nums) {
			res[i+1] = nums[start]
		}

		index -= 1
		start -= 1
		i += 2
	}

	for i := range nums {
		nums[i] = res[i]
	}
}

// @lc code=end

func Test324() {
	nums := []int{1, 1, 2, 1, 2, 2, 1}
	wiggleSort(nums)
	// [1,2,1,2,1,2,1]
	fmt.Println(nums)

	nums = []int{1, 5, 1, 1, 6, 4}
	// [1,6,1,5,1,4]
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{1, 3, 2, 2, 3, 1}
	// [2,3,1,3,1,2]
	wiggleSort(nums)
	fmt.Println(nums)

	nums = []int{5, 4, 5, 6}
	// [2,3,1,3,1,2]
	wiggleSort(nums)
	fmt.Println(nums)
}
