package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start

type pair373 struct {
	first  int
	second int
	sum    int
}

type pairs373 struct {
	pairs []*pair373
	sets  map[[2]int]struct{}
	nums1 []int
	nums2 []int
}

func (f *pairs373) diffuse(first, second int) {
	if first+1 < len(f.nums1) {
		f.add(first+1, second)
	}

	if second+1 < len(f.nums2) {
		f.add(first, second+1)
	}
}

func (f *pairs373) add(first, second int) {
	if first < len(f.nums1) && second < len(f.nums2) {
		if _, ok := f.sets[[2]int{first, second}]; ok {
			return
		}

		heap.Push(f, &pair373{
			first:  first,
			second: second,
			sum:    f.nums1[first] + f.nums2[second],
		})
		f.sets[[2]int{first, second}] = struct{}{}
	}
}

func (f *pairs373) Push(x any) {
	item := x.(*pair373)
	f.pairs = append(f.pairs, item)
}

func (f *pairs373) Pop() any {
	res := f.pairs[f.Len()-1]
	f.pairs = f.pairs[:f.Len()-1]
	return res
}

func (f pairs373) Len() int {
	return len(f.pairs)
}

func (f pairs373) Less(i, j int) bool {
	return f.pairs[i].sum < f.pairs[j].sum
}

func (f pairs373) Swap(i, j int) {
	f.pairs[i], f.pairs[j] = f.pairs[j], f.pairs[i]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &pairs373{
		pairs: make([]*pair373, 0, k),
		sets:  make(map[[2]int]struct{}),
		nums1: nums1,
		nums2: nums2,
	}

	h.add(0, 0)
	var res [][]int
	for i := 0; i < k; i++ {
		if h.Len() == 0 {
			break
		}

		pair := heap.Pop(h).(*pair373)
		res = append(res, []int{nums1[pair.first], nums2[pair.second]})
		h.diffuse(pair.first, pair.second)
	}

	return res
}

// @lc code=end

func Test373() {
	// [1,2],[1,4],[1,6]
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))

	// [1,1],[1,1]
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))

	// [1,3],[2,3]
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{3}, 3))

	// [[1,1],[1,1],[2,1],[1,2],[1,2],[2,2],[1,3],[1,3],[2,3]]
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10))
}
