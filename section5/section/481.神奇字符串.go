package section

import "fmt"

/*
 * @lc app=leetcode.cn id=481 lang=golang
 *
 * [481] 神奇字符串
 */

// @lc code=start
func magicalString(n int) int {
	magics := make([]int, 0, n)
	magics = append(magics, 1)

	index, count, num, ans := 1, 2, 2, 1
	for len(magics) < n {
		magics = append(magics, num)
		if num == 1 {
			ans += 1
		}

		if len(magics) == n {
			break
		}

		if count == 2 {
			magics = append(magics, num)
			if num == 1 {
				ans += 1
			}
		}

		index += 1
		count = magics[index]
		if num == 2 {
			num = 1
		} else {
			num = 2
		}
	}

	return ans
}

// @lc code=end

func Test481() {
	fmt.Println(magicalString(6) == 3)
	fmt.Println(magicalString(1) == 1)
}
