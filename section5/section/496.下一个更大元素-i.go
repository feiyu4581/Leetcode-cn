package section

import "fmt"

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		index := 0
		findPosition := false
		for index < len(nums2) {
			if findPosition && nums2[index] > num {
				break
			}

			if nums2[index] == num {
				findPosition = true
			}

			index += 1
		}

		if index < len(nums2) {
			ans = append(ans, nums2[index])
		} else {
			ans = append(ans, -1)
		}
	}

	return ans
}

// @lc code=end

func Test496() {
	// [-1,3,-1]
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	// [3,-1]
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
