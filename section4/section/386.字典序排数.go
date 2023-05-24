package section

import "fmt"

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, n)

	current := 1
	for i := 0; i < n; i++ {
		res[i] = current

		if current*10 <= n {
			current = current * 10
		} else if current%10 == 9 || current == n {
			for current%10 == 9 || current == n {
				current = current / 10
			}

			current += 1
		} else {
			current += 1
		}

	}

	return res
}

// @lc code=end

func Test386() {
	// [1,10,11,12,13,2,3,4,5,6,7,8,9]
	fmt.Println(lexicalOrder(13))
	// [1,2]
	fmt.Println(lexicalOrder(2))
	fmt.Println(lexicalOrder(111))
}
