package section

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 */

// @lc code=start
type pair struct {
	X int
	H int
}

type pairHeap []pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].H > h[j].H
}
func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	h := &pairHeap{}
	heap.Init(h)

	positions := make([]int, 0)
	for _, building := range buildings {
		positions = append(positions, building[0], building[1])
	}

	sort.Ints(positions)

	current := 0
	nums := make([][]int, 0)
	for _, position := range positions {
		for current < len(buildings) && buildings[current][0] <= position {
			heap.Push(h, pair{X: buildings[current][1], H: buildings[current][2]})
			current += 1
		}

		for h.Len() > 0 && (*h)[0].X <= position {
			heap.Pop(h)
		}

		max := 0
		if h.Len() > 0 {
			max = (*h)[0].H
		}
		if len(nums) == 0 || nums[len(nums)-1][1] != max {
			nums = append(nums, []int{position, max})
		}
	}

	return nums
}

func Test218() {
	// [[2,10],[10,25],[17,14],[20,0]]
	fmt.Println(getSkyline([][]int{
		{2, 13, 10},
		{10, 17, 25},
		{12, 20, 14},
	}))

	// [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	fmt.Println(getSkyline([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))
	// [[0,3],[5,0]]
	fmt.Println(getSkyline([][]int{
		{0, 2, 3},
		{2, 5, 3},
	}))

	// [[0,2147483647],[2147483647,0]]
	fmt.Println(getSkyline([][]int{
		{0, 2147483647, 2147483647},
	}))
}

// @lc code=end
