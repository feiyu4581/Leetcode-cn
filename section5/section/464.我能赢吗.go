package section

import "fmt"

/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 */

// @lc code=start

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	var canIWinHelper func(desiredTotal int, current int) bool

	cache := make(map[[2]int]bool)
	canIWinHelper = func(desiredTotal int, current int) (res bool) {
		cacheKey := [2]int{desiredTotal, current}
		if ans, ok := cache[cacheKey]; ok {
			return ans
		}

		defer func() {
			cache[cacheKey] = res
		}()

		for i := 1; i <= maxChoosableInteger; i++ {
			if current&(1<<i) == 0 {
				if i >= desiredTotal {
					return true
				}

				if !canIWinHelper(desiredTotal-i, current|(1<<i)) {
					return true
				}
			}
		}

		return false
	}

	return ((1+maxChoosableInteger)*maxChoosableInteger/2) >= desiredTotal && canIWinHelper(desiredTotal, 0)
}

// @lc code=end

func Test464() {
	fmt.Println(canIWin(5, 50) == false)
	fmt.Println(canIWin(20, 210) == false)
	fmt.Println(canIWin(10, 11) == false)
	fmt.Println(canIWin(10, 0) == true)
	fmt.Println(canIWin(10, 1) == true)
}
