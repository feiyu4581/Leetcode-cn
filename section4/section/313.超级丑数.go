package section

import "fmt"

/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	choose := make([]int, len(primes))
	for i := range choose {
		choose[i] = 1
	}

	dp[1] = 1
	for i := 2; i <= n; i++ {
		minValue := dp[choose[0]] * primes[0]
		for j := 1; j < len(primes); j++ {
			if value := dp[choose[j]] * primes[j]; value < minValue {
				minValue = value
			}
		}

		for j := 0; j < len(primes); j++ {
			if minValue == dp[choose[j]]*primes[j] {
				choose[j] = choose[j] + 1
			}
		}
		dp[i] = minValue
	}

	return dp[n]
}

// @lc code=end

func Test313() {
	// [1,2,4,7,8,13,14,16,19,26,28,32]
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}) == 32)
	fmt.Println(nthSuperUglyNumber(1, []int{2, 3, 5}) == 1)
}
