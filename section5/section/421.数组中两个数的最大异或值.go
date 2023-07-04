package section

import "fmt"

/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	res := 0
	for k := 30; k >= 0; k-- {
		seen := make(map[int]struct{}, len(nums))
		for _, num := range nums {
			seen[num>>k] = struct{}{}
		}

		res = res*2 + 1
		findNum := false
		for _, num := range nums {
			if _, ok := seen[res^(num>>k)]; ok {
				findNum = true
				break

			}
		}

		if !findNum {
			res -= 1
		}
	}

	return res
}

// @lc code=end

func Test421() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}) == 28)
	fmt.Println(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}) == 127)
}
