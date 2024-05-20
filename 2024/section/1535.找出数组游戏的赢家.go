package section

import "fmt"

/*
 * @lc app=leetcode.cn id=1535 lang=golang
 *
 * [1535] 找出数组游戏的赢家
 */

// @lc code=start
func getWinner(arr []int, k int) int {
	count, current := 0, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > current {
			current = arr[i]
			count = 1
		} else {
			count += 1
		}

		if count == k {
			return current
		}
	}

	// 当遍历一次之后，留下来的数据就一定是最大值，对于最大值来说，不管需要比较多少次结果都是一样的
	return current
}

// @lc code=end

func Test1535() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2) == 5)
	fmt.Println(getWinner([]int{3, 2, 1}, 10) == 3)
	fmt.Println(getWinner([]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7) == 9)
	fmt.Println(getWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000) == 99)
}
