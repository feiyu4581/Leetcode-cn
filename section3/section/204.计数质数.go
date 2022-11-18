package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

func Test204() {
	fmt.Println(countPrimes(10) == 4)
	fmt.Println(countPrimes(0) == 0)
	fmt.Println(countPrimes(1) == 0)
}

// @lc code=end
