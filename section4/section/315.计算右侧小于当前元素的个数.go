package section

import "fmt"

/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))

	var merge func([]int) []int
	merge = func(positions []int) []int {
		if len(positions) <= 1 {
			return positions
		}

		mid := len(positions) / 2
		left := merge(positions[:mid])
		right := merge(positions[mid:])

		i, j := 0, 0
		newPositions := make([]int, len(positions))
		for i < len(left) && j < len(right) {
			if nums[left[i]] <= nums[right[j]] {
				res[left[i]] += j
				newPositions[i+j] = left[i]
				i += 1
			} else {
				newPositions[i+j] = right[j]
				j += 1
			}
		}

		for i < len(left) {
			res[left[i]] += j
			newPositions[i+j] = left[i]
			i += 1
		}

		for j < len(right) {
			newPositions[i+j] = right[j]
			j += 1
		}

		return newPositions
	}

	positions := make([]int, len(nums))
	for i := range nums {
		positions[i] = i
	}
	merge(positions)
	return res
}

// @lc code=end

func Test315() {
	// [2,1,1,0]
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))

	// [0]
	fmt.Println(countSmaller([]int{-1}))

	// [0, 0]
	fmt.Println(countSmaller([]int{-1, -1}))
}
