package section

import "fmt"

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	nums := make([][]int, len(s)+1)
	for index := range nums {
		nums[index] = make([]int, len(t)+1)
		nums[index][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i && j <= len(t); j++ {
			nums[i][j] = nums[i-1][j]
			if s[i-1] == t[j-1] {
				nums[i][j] += nums[i-1][j-1]
			}
			fmt.Printf("[%d, %d]=[%s, %s]=%d\n", i, j, s[0:i], t[0:j], nums[i][j])
		}
	}

	return nums[len(s)][len(t)]
}

func Test115() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("bab", "b"))
	fmt.Println(numDistinct("babgbag", "bag"))
}

// @lc code=end
