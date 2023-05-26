package section

import "fmt"

/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */

// @lc code=start
func lastRemaining(n int) int {
	flags := make([]bool, 0)
	compute := func(num int) int {
		for i := len(flags) - 1; i >= 0; i-- {
			if flags[i] {
				num = 2 * num
			} else {
				num = 2*num - 1
			}
		}

		return num
	}

	reverse := false
	current := n
	for current > 1 {
		if reverse && current%2 == 0 {
			flags = append(flags, false)
		} else {
			flags = append(flags, true)
		}

		current = current / 2
		reverse = !reverse
	}

	return compute(1)
}

// @lc code=end

func Test390() {
	fmt.Println(lastRemaining(4) == 2)
	fmt.Println(lastRemaining(153) == 94)
	fmt.Println(lastRemaining(9) == 6)
	fmt.Println(lastRemaining(1) == 1)
}
