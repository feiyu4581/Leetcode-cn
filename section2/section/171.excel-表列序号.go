package section

import "fmt"

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	current := 0
	for _, title := range columnTitle {
		num := title - 'A' + 1
		current = current * 26 + int(num)
	}

	return current
}

func Test171() {
	fmt.Println(titleToNumber("A") == 1)
	fmt.Println(titleToNumber("AZ") == 52)
	fmt.Println(titleToNumber("YZ") == 676)
	fmt.Println(titleToNumber("FXSHRXW") == 2147483647)
	fmt.Println(titleToNumber("AB") == 28)
	fmt.Println(titleToNumber("ZY") == 701)
}
// @lc code=end

