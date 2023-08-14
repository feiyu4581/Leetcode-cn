package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=458 lang=golang
 *
 * [458] 可怜的小猪
 */

// @lc code=start
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets))/math.Log(float64(states)) - 1e-5))
}

// @lc code=end

func Test458() {
	fmt.Println(poorPigs(1000, 15, 60) == 5)
	fmt.Println(poorPigs(4, 15, 15) == 2)
	fmt.Println(poorPigs(4, 15, 30) == 2)
}
