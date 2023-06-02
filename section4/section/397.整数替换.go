package section

import "fmt"

/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start

func min397(left, right int) int {
	if left < right {
		return left
	}

	return right
}

func integerReplacementWithCache(n int, cache map[int]int) int {
	if count, ok := cache[n]; ok {
		return count
	}

	count := 1
	if n == 1 {
		return count
	}

	if n%2 == 0 {
		count += integerReplacementWithCache(n/2, cache)
	} else {
		count += min397(integerReplacementWithCache(n-1, cache), integerReplacementWithCache(n+1, cache))
	}

	cache[n] = count
	return count
}

func integerReplacement(n int) int {
	return integerReplacementWithCache(n, map[int]int{}) - 1
}

// @lc code=end

func Test397() {
	fmt.Println(integerReplacement(8) == 3)
	fmt.Println(integerReplacement(7) == 4)
}
