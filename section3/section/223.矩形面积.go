package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	min := func(left, right int) int {
		if left < right {
			return left
		}

		return right
	}

	if ax1 > bx1 {
		ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 = bx1, by1, bx2, by2, ax1, ay1, ax2, ay2
	}

	mid := 0
	if bx1 < ax2 {
		x := min(ax2, bx2) - bx1
		y1, y2, y3, y4 := ay1, ay2, by1, by2
		if y1 > y3 {
			y1, y2, y3, y4 = y3, y4, y1, y2
		}

		y := 0
		if y2 >= y3 {
			y = min(y2, y4) - y3
		}

		mid = x * y
	}

	left := (ax2 - ax1) * (ay2 - ay1)
	right := (bx2 - bx1) * (by2 - by1)

	//fmt.Printf("%d-%d-%d\n", left, right, mid)
	return left + right - mid
}

func Test223() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2) == 45)
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2) == 16)
	fmt.Println(computeArea(-2, -2, 2, 2, -1, 4, 1, 6) == 20)
}

// @lc code=end
