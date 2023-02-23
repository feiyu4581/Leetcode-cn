package section

import "fmt"

/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 多数元素 II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	threshold := float64(len(nums)) / 3.0
	res := make([]int, 0)

	counter := make(map[int]float64, 0)
	for _, num := range nums {
		if _, ok := counter[num]; !ok {
			counter[num] = 0
		}

		counter[num] += 1
		if counter[num] > threshold {
			res = append(res, num)
			counter[num] = float64(-len(nums))
		}
	}

	return res
}

// @lc code=end

func Test229() {
	// [3]
	fmt.Println(majorityElement([]int{3, 2, 3}))
	// [1]
	fmt.Println(majorityElement([]int{1}))
	// [1, 2]
	fmt.Println(majorityElement([]int{1, 2}))
}
