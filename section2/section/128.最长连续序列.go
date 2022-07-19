package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start

func computeIndexNumMaps(num int, numMaps map[int]struct{}, indexMaps map[int]int, left bool) int {
	if _, ok := numMaps[num]; !ok {
		return 0
	}

	if index, ok := indexMaps[num]; ok {
		return index
	}

	if left {
		indexMaps[num] = computeIndexNumMaps(num-1, numMaps, indexMaps, left) + 1
	} else {
		indexMaps[num] = computeIndexNumMaps(num+1, numMaps, indexMaps, left) + 1
	}
	return indexMaps[num]
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numMaps := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numMaps[num] = struct{}{}
	}

	leftNumMaps := make(map[int]int, len(nums))
	rightNumMaps := make(map[int]int, len(nums))

	maxLength := math.MinInt
	for _, num := range nums {
		index := computeIndexNumMaps(num, numMaps, leftNumMaps, true) + computeIndexNumMaps(num, numMaps, rightNumMaps, false) - 1
		if index > maxLength {
			maxLength = index
		}
	}

	return maxLength
}

func Test128() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}) == 4)
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}) == 9)
}

// @lc code=end
