package section

import "fmt"

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	nums := make([]int, 0, 32)
	res := make([]int, n+1)
	res[0] = 0
	nums = append(nums, 0)

	count := 0
	for i := 1; i <= n; i++ {
		carry := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == 0 {
				nums[j] += 1
				count += 1
				carry = 0
				break
			} else {
				nums[j] = 0
				count -= 1
				carry = 1
			}
		}

		if carry == 1 {
			count += 1
			nums = append(nums, carry)
		}

		res[i] = count
	}

	return res
}

// @lc code=end

func Test338() {
	// [0,1,1,2,1,2,2,3,1]
	fmt.Println(countBits(8))
	// [0,1,1]
	fmt.Println(countBits(2))
	// [0,1,1,2,1,2]
	fmt.Println(countBits(5))
}
