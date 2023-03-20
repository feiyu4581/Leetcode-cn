package section

import "fmt"

/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end

func Test292() {
	fmt.Println(canWinNim(4) == false)
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(2))
}
