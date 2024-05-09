package section

import "fmt"

/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start

const tilingMod = 1e9 + 7

/*
 * 通用的块组织方式分成以下四种，因此，可以通过 dp 来解决问题，低于 3 的情况下，直接枚举即可
 * 超过 3 的情况下，dp[n] = dp[n - 3] + dp[n - 3] + dp[n - 2] + dp[n - 1], 分别表示开头使用以下四块方式
 * + + -     - - +     + +     +
 * + - -     - + +     + +     +
 * 超过 4 的情况下，存在一种特殊情况
 * + + - - +
 * + - - + +，此时将两个方块组合起来，左右两个存在上下位置的差异，因此存在两种情况，中间的两个方块存在 n - 4 种情况（长度）
 * 因此 dp[n] += 2 * (dp[n - 4] + dp[n - 5] + ... + dp[1])
 */
func numTilings(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 5
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2], dp[3] = 1, 1, 2, 5

	offset := 0
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-3]*2 + dp[i-2] + dp[i-1]

		// 将 for 循环的结果进行优化，减少循环次数
		offset += dp[i-4] * 2
		dp[i] += offset

		//for j := i - 4; j >= 0; j-- {
		//	dp[i] += dp[j] * 2
		//}

		if dp[i] >= tilingMod {
			dp[i] = dp[i] % tilingMod
		}
	}

	return dp[n]
}

// @lc code=end

func Test790() {
	fmt.Println(numTilings(3) == 5)
	fmt.Println(numTilings(1) == 1)
	fmt.Println(numTilings(4) == 11) // 11
	fmt.Println(numTilings(6) == 53) // 53
	fmt.Println(numTilings(743) == 561645440)
}
