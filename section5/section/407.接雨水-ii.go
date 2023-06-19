package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=407 lang=golang
 *
 * [407] 接雨水 II
 */

// @lc code=start

func min407(left, right int) int {
	if left < right {
		return left
	}

	return right
}

type box407 struct {
	row    int
	column int

	water int
}

type boxs407 []box407

func (w boxs407) Len() int {
	return len(w)
}

func (w boxs407) Less(i, j int) bool {
	return w[i].water < w[j].water
}

func (w boxs407) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *boxs407) Push(x any) {
	*w = append(*w, x.(box407))
}

func (w *boxs407) Pop() any {
	old := *w
	n := len(old)
	x := old[n-1]
	*w = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return 0
	}

	b := &boxs407{}
	waters := make([][]int, m)
	visits := make([][]bool, m)

	for i := 0; i < m; i++ {
		waters[i] = make([]int, n)
		visits[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		heap.Push(b, box407{
			row:    i,
			column: 0,
			water:  heightMap[i][0],
		})

		heap.Push(b, box407{
			row:    i,
			column: n - 1,
			water:  heightMap[i][n-1],
		})
	}

	for j := 0; j < n; j++ {
		heap.Push(b, box407{
			row:    0,
			column: j,
			water:  heightMap[0][j],
		})
		heap.Push(b, box407{
			row:    m - 1,
			column: j,
			water:  heightMap[m-1][j],
		})
	}

	dirs := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for b.Len() > 0 {
		box := heap.Pop(b).(box407)
		for _, dir := range dirs {
			x, y := box.row+dir[0], box.column+dir[1]
			if x <= 0 || x >= m-1 || y <= 0 || y >= n-1 || visits[x][y] {
				continue
			}

			if waters[x][y] == 0 || waters[x][y] > box.water {
				water := box.water
				if heightMap[x][y] > box.water {
					water = heightMap[x][y]
				}

				waters[x][y] = water
				heap.Push(b, box407{
					row:    x,
					column: y,
					water:  water,
				})
			}

			visits[x][y] = true
		}
	}

	res := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			res += waters[i][j] - heightMap[i][j]
		}
	}

	return res
}

// @lc code=end

func Test407() {
	// 25
	fmt.Println(trapRainWater([][]int{
		{14, 17, 18, 16, 14, 16},
		{17, 3, 10, 2, 3, 8},
		{11, 10, 4, 7, 1, 7},
		{13, 7, 2, 9, 8, 10},
		{13, 1, 3, 4, 8, 6},
		{20, 3, 3, 9, 10, 8},
	}) == 25)

	fmt.Println(trapRainWater([][]int{
		{5, 8, 7, 7},
		{5, 2, 1, 5},
		{7, 1, 7, 1},
		{8, 9, 6, 9},
		{9, 8, 9, 9},
	}) == 12)

	fmt.Println(trapRainWater([][]int{
		{12, 13, 1, 12},
		{13, 4, 13, 12},
		{13, 8, 10, 12},
		{12, 13, 12, 12},
		{13, 13, 13, 13},
	}) == 14)

	fmt.Println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}) == 4)

	fmt.Println(trapRainWater([][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}) == 10)
}
