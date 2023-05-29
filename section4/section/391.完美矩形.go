package section

import "fmt"

/*
 * @lc app=leetcode.cn id=391 lang=golang
 *
 * [391] 完美矩形
 */

// @lc code=start

func isRectangleCover(rectangles [][]int) bool {
	type point struct{ x, y int }
	area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	cnt := map[point]int{}
	for _, rect := range rectangles {
		x, y, a, b := rect[0], rect[1], rect[2], rect[3]
		area += (a - x) * (b - y)

		minX = min391(minX, x)
		minY = min391(minY, y)
		maxX = max391(maxX, a)
		maxY = max391(maxY, b)

		cnt[point{x, y}]++
		cnt[point{x, b}]++
		cnt[point{a, y}]++
		cnt[point{a, b}]++
	}

	if area != (maxX-minX)*(maxY-minY) || cnt[point{minX, minY}] != 1 || cnt[point{minX, maxY}] != 1 || cnt[point{maxX, minY}] != 1 || cnt[point{maxX, maxY}] != 1 {
		return false
	}

	delete(cnt, point{minX, minY})
	delete(cnt, point{minX, maxY})
	delete(cnt, point{maxX, minY})
	delete(cnt, point{maxX, maxY})

	for _, c := range cnt {
		if c != 2 && c != 4 {
			return false
		}
	}
	return true
}

func min391(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max391(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

func Test391() {
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}))

	fmt.Println(isRectangleCover([][]int{
		{1, 1, 2, 3},
		{1, 3, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
	}) == false)

	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4},
	}) == false)
}
