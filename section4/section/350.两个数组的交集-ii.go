package section

import "fmt"

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect350(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect350(nums2, nums1)
	}

	counter := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		counter[num] += 1
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if count := counter[num]; count > 0 {
			res = append(res, num)
			counter[num] -= 1
		}
	}

	return res
}

// @lc code=end

func Test350() {
	// [2,2]
	fmt.Println(intersect350([]int{1, 2, 2, 1}, []int{2, 2}))

	// [4,9]
	fmt.Println(intersect350([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
