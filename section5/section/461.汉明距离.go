package section

import "fmt"

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y

	ans := 0
	for i := 0; i <= 31; i++ {
		if z&(1<<i) != 0 {
			ans++
		}
	}

	return ans
}

// @lc code=end

func Test461() {
	fmt.Println(hammingDistance(1, 4) == 2)

}
