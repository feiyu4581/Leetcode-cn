package section

import "fmt"

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countMaps := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			countMaps[num1+num2]++
		}
	}

	ans := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			ans += countMaps[-num3-num4]
		}
	}

	return ans
}

// @lc code=end

func Test454() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}) == 2)
	fmt.Println(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}) == 1)
}
