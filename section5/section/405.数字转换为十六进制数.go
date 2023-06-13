package section

import "fmt"

/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start

func convertNumToByte(num int) byte {
	if num >= 10 {
		return 'a' + byte(num) - 10
	} else {
		return '0' + byte(num)
	}
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hex := make([]int, 8)
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	for i := 7; i >= 0 && num > 0; i-- {
		hex[i] = num % 16

		num = num / 16
	}

	if isNegative {
		for i := 0; i < 8; i++ {
			hex[i] = 15 - hex[i]
		}

		for i := 7; i >= 0; i-- {
			if hex[i] == 15 {
				hex[i] = 0
			} else {
				hex[i] += 1
				break
			}
		}
	}

	res := make([]byte, 0, 8)
	for i := 0; i < 8; i++ {
		if hex[i] == 0 && len(res) == 0 {
			continue
		}

		res = append(res, convertNumToByte(hex[i]))
	}
	return string(res)
}

// @lc code=end

func Test405() {
	fmt.Println(toHex(26) == "1a")
	fmt.Println(toHex(-1) == "ffffffff")
}
