package section

import "fmt"

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0

	current := uint32(1)
	index := 0
	for num >= current && index <= 31 {
		if num&current > 0 {
			count += 1
		}

		current = current << 1
		index += 1
	}

	return count
}

func Test191() {
	fmt.Println(hammingWeight(0b1011) == 3)
	fmt.Println(hammingWeight(0b10000000) == 1)
	fmt.Println(hammingWeight(0b11111111111111111111111111111101) == 31)
}

// @lc code=end
