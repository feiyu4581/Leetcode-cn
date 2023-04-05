package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}

// @lc code=end

func Test319() {
	fmt.Println(bulbSwitch(6) == 2)
	fmt.Println(bulbSwitch(3) == 1)
	fmt.Println(bulbSwitch(0) == 0)
	fmt.Println(bulbSwitch(1) == 1)
}
