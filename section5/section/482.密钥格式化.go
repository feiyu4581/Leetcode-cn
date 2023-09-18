package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1))
	group := len(s) / k
	firstGroup := len(s) % k

	ans := make([]byte, 0, len(s)+group)
	missingNums := firstGroup
	for i := 0; i < len(s); i++ {
		if missingNums == 0 {
			if len(ans) > 0 {
				ans = append(ans, '-')
			}
			missingNums = k
		}

		ans = append(ans, s[i])
		missingNums -= 1
	}

	return string(ans)
}

// @lc code=end

func Test482() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4) == "5F3Z-2E9W")
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2) == "2-5G-3J")
}
