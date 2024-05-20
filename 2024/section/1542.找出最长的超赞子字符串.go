package section

import "fmt"

/*
 * @lc app=leetcode.cn id=1542 lang=golang
 *
 * [1542] 找出最长的超赞子字符串
 */

// @lc code=start
func longestAwesome(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	prefix := map[int]int{0: -1}
	ans := 0
	sequence := 0
	for j := 0; j < len(s); j++ {
		digit := int(s[j] - '0')
		// 这里的 sequence 代表了[0:j] 的所有数字的奇偶性
		// 因为只有 0-9 十个数字，所以可以用一个 int 的二进制来表示，0～9每一位代表一个数字的奇偶性
		// 0 代表偶数，1 代表奇数，因此每多便利一个数字，只要异或上这个数字的二进制位，就可以得到这个数字的奇偶性
		// f(0:j) = f(0:j-1) ^ (1 << digit)
		sequence ^= 1 << digit
		// prefix 代表了每一个 sequence 第一次出现的位置，如果二者相同，那么代表着 f(0:j) 和 f(0:prevIndex) 存在相同的奇偶性
		// 那么 f(0:j) - f(0:prevIndex) = f(prevIndex+1:j)，此时 f(prevIndex+1:j) 的的所有数字一定都是偶数（奇数 - 偶数 = 偶数，偶数 - 偶数 = 偶数）
		if prevIndex, ok := prefix[sequence]; ok {
			ans = max(ans, j-prevIndex)
		} else {
			prefix[sequence] = j
		}
		for k := 0; k < 10; k++ {
			// 从 0 ～ 10，将 sequence 微调一位，此时依然满足的情况下，表示 f(prevIndex+1;j) 除了微调的哪一位全部都是偶数（微调的是奇数）
			// 此时只有一个奇数的情况下，f(prevIndex+1:j) 是一个超赞字符串
			if prevIndex, ok := prefix[sequence^(1<<k)]; ok {
				ans = max(ans, j-prevIndex)
			}
		}
	}
	return ans
}

// @lc code=end

func Test1542() {
	fmt.Println(longestAwesome("3242415") == 5)
	fmt.Println(longestAwesome("12345678") == 1)
	fmt.Println(longestAwesome("213123") == 6)
	fmt.Println(longestAwesome("00") == 2)
	fmt.Println(longestAwesome("9498331") == 3)
}
