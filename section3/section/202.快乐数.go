package section

import "fmt"

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start

func recursiveIsHappy(n int, cache map[int]struct{}) bool {
	if n == 1 {
		return true
	}

	if _, ok := cache[n]; ok {
		return false
	}

	value, num := 0, n
	for num > 0 {
		residue := num % 10
		value += residue * residue

		num = num / 10
	}

	cache[n] = struct{}{}
	return recursiveIsHappy(value, cache)
}

func isHappy(n int) bool {
	return recursiveIsHappy(n, make(map[int]struct{}, 0))
}

func Test202() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2) == false)
}

// @lc code=end
