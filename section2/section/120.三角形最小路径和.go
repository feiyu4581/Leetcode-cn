package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start

func minimumDepthTotal(triangle [][]int, depth int, column int, cache map[string]int) int {
	if depth >= len(triangle) || column >= len(triangle[depth]) {
		return 0
	}

	key := fmt.Sprintf("%d-%d", depth, column)
	if num, ok := cache[key]; ok {
		return num
	}

	left := minimumDepthTotal(triangle, depth+1, column, cache)
	right := minimumDepthTotal(triangle, depth+1, column+1, cache)
	if right < left {
		left = right
	}

	cache[key] = left + triangle[depth][column]
	return cache[key]
}

func minimumTotal2(triangle [][]int) int {
	return minimumDepthTotal(triangle, 0, 0, make(map[string]int, 0))
}

func minimumTotal(triangle [][]int) int {
	for row := range triangle {
		if row == 0 {
			continue
		}
		for column := range triangle[row] {
			value := math.MaxInt
			if column > 0 {
				value = triangle[row-1][column-1]
			}

			if column < len(triangle[row-1]) && value > triangle[row-1][column] {
				value = triangle[row-1][column]
			}

			triangle[row][column] += value
		}
	}

	minValue := math.MaxInt
	for _, value := range triangle[len(triangle)-1] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func Test120() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}) == 11)
	fmt.Println(minimumTotal([][]int{{-10}}) == -10)
	fmt.Println(minimumTotal([][]int{{-1}, {-2, -3}}) == -4)
}

// @lc code=end
