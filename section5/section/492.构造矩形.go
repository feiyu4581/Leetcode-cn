package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	for i := sqrt; i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}

	return []int{area, 1}
}

// @lc code=end

func Test492() {
	//  [2, 2]
	fmt.Println(constructRectangle(4))
	// [37, 1]
	fmt.Println(constructRectangle(37))
	// [427,286]
	fmt.Println(constructRectangle(122122))
}
