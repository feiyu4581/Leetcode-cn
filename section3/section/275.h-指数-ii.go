package section

import "fmt"

/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex2(citations []int) int {
	left, right := 0, len(citations)-1

	for left <= right {
		mid := left + (right-left)>>1
		if citations[mid] >= len(citations)-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return len(citations) - left
}

// @lc code=end

func Test275() {
	fmt.Println(hIndex2([]int{0}) == 0)
	fmt.Println(hIndex2([]int{0, 1, 3, 5, 6}) == 3)
	fmt.Println(hIndex2([]int{1, 2, 100}) == 2)
}
