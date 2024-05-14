package section

import "fmt"

/*
 * @lc app=leetcode.cn id=2244 lang=golang
 *
 * [2244] 完成所有任务需要的最少轮数
 */

// @lc code=start
func minimumRounds(tasks []int) int {
	counter := make(map[int]int)
	for _, task := range tasks {
		counter[task] += 1
	}

	res := 0
	for _, num := range counter {
		if num == 1 {
			return -1
		}

		res += num / 3
		// 如果 num 不能被 3 整除，根据余数的结果如下
		// 2: 此时直接加 1 即可，满足一轮完成的条件
		// 1: 此时，num 要么是 4，要么是 4 + 3*n，因此可以转换为 (n - 1) * 3 + 2 * 2，此时也只需要加 1 即可
		if num%3 != 0 {
			res += 1
		}
	}

	return res
}

// @lc code=end

func Test2244() {
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}) == 4)
	fmt.Println(minimumRounds([]int{2, 3, 3}) == -1)
}
