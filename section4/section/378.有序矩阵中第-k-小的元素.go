package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start

type pair378 struct {
	x     int
	y     int
	value int
}

type pairs378 struct {
	pairs  []*pair378
	cache  map[[2]int]struct{}
	matrix [][]int
}

func (p *pairs378) Add(x, y int) {
	if x >= len(p.matrix) || y >= len(p.matrix[0]) {
		return
	}

	if _, ok := p.cache[[2]int{x, y}]; ok {
		return
	}

	heap.Push(p, &pair378{
		x:     x,
		y:     y,
		value: p.matrix[x][y],
	})

	p.cache[[2]int{x, y}] = struct{}{}
}

func (p *pairs378) Push(x any) {
	item := x.(*pair378)
	p.pairs = append(p.pairs, item)
}

func (p *pairs378) Pop() any {
	res := p.pairs[p.Len()-1]
	p.pairs = p.pairs[:p.Len()-1]
	return res
}

func (p pairs378) Len() int {
	return len(p.pairs)
}

func (p pairs378) Less(i, j int) bool {
	return p.pairs[i].value < p.pairs[j].value
}

func (p pairs378) Swap(i, j int) {
	p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i]
}

func kthSmallest(matrix [][]int, k int) int {
	h := &pairs378{
		pairs:  make([]*pair378, 0, k),
		cache:  make(map[[2]int]struct{}, k),
		matrix: matrix,
	}

	h.Add(0, 0)
	res := matrix[0][0]
	for i := 0; i < k; i++ {
		item := heap.Pop(h).(*pair378)
		h.Add(item.x+1, item.y)
		h.Add(item.x, item.y+1)

		res = item.value
	}

	return res
}

// @lc code=end

func Test378() {
	fmt.Println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8) == 13)

	fmt.Println(kthSmallest([][]int{
		{1, 3, 5},
		{6, 7, 12},
		{11, 14, 14},
	}, 6) == 11)
}
