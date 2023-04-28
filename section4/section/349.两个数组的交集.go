package section

import "fmt"

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		counter[num] = 0
	}

	for _, num := range nums2 {
		if _, ok := counter[num]; ok {
			counter[num] = 1
		}
	}

	res := make([]int, 0)
	for num, count := range counter {
		if count == 1 {
			res = append(res, num)
		}
	}

	return res
}

// @lc code=end

func Test349() {
	// [2]
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))

	// [9, 4] or [4, 9]
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
