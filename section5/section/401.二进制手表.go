package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start

func chooseBinaryNums(num int, limit int, turnedOn int) []int {
	if turnedOn == 0 {
		return []int{0}
	}

	res := make([]int, 0)

	var dfs func(int, int, int, int)
	dfs = func(start, left, currentNum, stepNum int) {
		if left == 0 {
			if currentNum <= limit {
				res = append(res, currentNum)
			}
		} else if start < num {
			dfs(start+1, left-1, currentNum+stepNum, stepNum*2)
			if num-start > left {
				dfs(start+1, left, currentNum, stepNum*2)
			}
		}
	}

	dfs(0, turnedOn, 0, 1)
	return res
}

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)
	for left := 0; left <= 3 && left <= turnedOn; left++ {
		right := turnedOn - left
		if right >= 6 {
			continue
		}

		for _, i := range chooseBinaryNums(4, 11, left) {
			for _, j := range chooseBinaryNums(6, 59, right) {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}

	return res
}

// @lc code=end

func Test401() {
	// ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
	fmt.Println(readBinaryWatch(1))
	// []
	fmt.Println(readBinaryWatch(9))
}
