package section

import "fmt"

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	nums := make([]uint32, 0)
	for num != 0 {
		nums = append(nums, num%2)
		num = num / 2
	}

	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = res * 2
		if i < len(nums) {
			res = res + nums[i]
		}
	}

	return res
}

func Test190() {
	fmt.Println(reverseBits(43261596) == 964176192)
	fmt.Println(reverseBits(4294967293) == 3221225471)
}

// @lc code=end
