package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

type item239 struct {
	Value int
	Index int
}

type h239 struct {
	items  []*item239
	indexs map[int]*item239
}

func (h *h239) Len() int {
	return len(h.items)
}

func (h *h239) Less(i, j int) bool {
	return h.items[i].Value > h.items[j].Value
}

func (h *h239) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].Index = i
	h.items[j].Index = j
}

func (h *h239) Push(x any) {
	h.items = append(h.items, x.(*item239))
}

func (h *h239) Pop() any {
	h.items = h.items[:len(h.items)-1]
	return nil
}

func (h *h239) PushNum(num, index int) {
	item := &item239{
		Value: num,
		Index: len(h.items),
	}

	h.indexs[index] = item
	heap.Push(h, item)
}

func (h *h239) ReplaceNum(source, target, num int) {
	item := h.indexs[source]
	delete(h.indexs, source)
	h.indexs[target] = item

	item.Value = num
	heap.Fix(h, item.Index)
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	h := &h239{
		items:  make([]*item239, 0),
		indexs: make(map[int]*item239, 0),
	}

	heap.Init(h)

	for i := 0; i < k && i < len(nums); i++ {
		h.PushNum(nums[i], i)
	}

	res := make([]int, 0)
	res = append(res, h.items[0].Value)

	for i := k; i < len(nums); i++ {
		h.ReplaceNum(i-k, i, nums[i])
		res = append(res, h.items[0].Value)
	}

	return res
}

// @lc code=end

func Test239() {
	// [3,3,5,5,6,7]
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// [1]
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}
