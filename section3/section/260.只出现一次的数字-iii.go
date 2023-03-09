package section

import "fmt"

/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	x := 0
	// 所有出现两次的数据，都会随着异或操作被抵消，所以最终的 x 一定等于两个只出现一次的异或结果
	for _, num := range nums {
		x = x ^ num
	}

	// 找到这个异或值出现的第一个 1，此时两个数字这个位置上一定有一个1，一个 0，此时将所有数据根据这个位置是0，还是 1进行分组
	// 此时就可以得到两个数组，每个数组里面的数字除了一个数字只出现一次以外，其他所有数字都会出现两次，此时简单异或所有数字即可
	lastX := x & -x
	left, right := 0, 0

	for _, num := range nums {
		if num&lastX > 0 {
			left = left ^ num
		} else {
			right = right ^ num
		}
	}

	return []int{left, right}
}

// @lc code=end

func Test260() {
	// [3,5]
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	// [-1,0]
	fmt.Println(singleNumber([]int{-1, 0}))
	// [0,1]
	fmt.Println(singleNumber([]int{0, 1}))
}
