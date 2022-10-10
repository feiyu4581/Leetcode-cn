package section

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})

	if nums[0] == 0 {
		return "0"
	}

	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func Test179() {
	fmt.Println(largestNumber([]int{10, 2}) == "210")
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}) == "9534330")
	fmt.Println(largestNumber([]int{111311, 1113}) == "1113111311")
	fmt.Println(largestNumber([]int{8308, 8308, 830}) == "83088308830")
}

// @lc code=end
