package section

import "fmt"

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	res := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for len(res) > 0 && res[len(res)-1] > num[i] && k > 0 {
			res = res[:len(res)-1]
			k -= 1
		}

		res = append(res, num[i])
	}

	if k > 0 {
		res = res[:len(res)-k]
	}

	for len(res) > 0 && res[0] == '0' {
		res = res[1:]
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}

// @lc code=end

func Test402() {
	fmt.Println(removeKdigits("9510982", 4))
	fmt.Println(removeKdigits("1432219", 3) == "1219")
	fmt.Println(removeKdigits("10200", 1) == "200")
	fmt.Println(removeKdigits("13200", 1) == "1200")
	fmt.Println(removeKdigits("10", 2) == "0")
}
