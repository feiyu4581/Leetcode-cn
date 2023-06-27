package section

import "fmt"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	carry := uint8(0)
	res := make([]byte, 0, len(num1)+len(num2))

	i := 0
	for i < len(num1) || i < len(num2) {
		left, right := byte('0'), byte('0')
		if i < len(num1) {
			left = num1[len(num1)-i-1]
		}

		if i < len(num2) {
			right = num2[len(num2)-i-1]
		}

		num := left + right + carry - '0'
		carry = 0
		if num > '9' {
			carry = 1
			num -= 10
		}

		res = append(res, num)
		i += 1
	}

	if carry > 0 {
		res = append(res, carry+'0')
	}

	i, j := 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i += 1
		j -= 1
	}

	return string(res)
}

// @lc code=end

func Test415() {
	fmt.Println(addStrings("11", "123") == "134")
	fmt.Println(addStrings("456", "77") == "533")
	fmt.Println(addStrings("0", "0") == "0")
	fmt.Println(addStrings("1", "9") == "10")
}
