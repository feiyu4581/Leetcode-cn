package section

import "fmt"

/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	ans := 0
	bit := 1
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			if num&bit != 0 {
				count += 1
			}
		}
		bit = bit << 1
		ans += count * (len(nums) - count)
	}

	return ans
}

// @lc code=end

func Test477() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}) == 6)
	fmt.Println(totalHammingDistance([]int{4, 14, 4}) == 4)
}
