package section

import (
	"container/heap"
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=480 lang=golang
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start

type medianItem struct {
	Value int
	Index int
}

type MediaHeap struct {
	Items    []*medianItem
	ItemMaps map[int][]*medianItem
	Reversed bool
}

func NewMediaHeap(isReversed bool) *MediaHeap {
	return &MediaHeap{
		ItemMaps: make(map[int][]*medianItem, 0),
		Reversed: isReversed,
	}
}

func (m *MediaHeap) Print() string {
	value := make([]string, 0)
	for _, item := range m.Items {
		value = append(value, fmt.Sprintf("%d", item.Value))
	}

	return strings.Join(value, ",")
}

func (m *MediaHeap) Len() int {
	return len(m.Items)
}

func (m *MediaHeap) Less(i, j int) bool {
	if m.Reversed {
		return m.Items[i].Value > m.Items[j].Value
	}

	return m.Items[i].Value < m.Items[j].Value
}

func (m *MediaHeap) Swap(i, j int) {
	m.Items[i], m.Items[j] = m.Items[j], m.Items[i]

	m.Items[i].Index = i
	m.Items[j].Index = j
}

func (m *MediaHeap) Push(x any) {
	item := &medianItem{
		Value: x.(int),
		Index: m.Len(),
	}

	m.ItemMaps[item.Value] = append(m.ItemMaps[item.Value], item)
	m.Items = append(m.Items, item)
}

func (m *MediaHeap) Pop() any {
	item := m.Items[m.Len()-1]
	m.Items = m.Items[:m.Len()-1]

	// 这里需要保证删除的时候，永远是删除最先插入的节点（如果值相同的话）
	m.ItemMaps[item.Value] = m.ItemMaps[item.Value][1:]
	return item.Value
}

type MedianHelper struct {
	minHeap *MediaHeap
	maxHeap *MediaHeap
	Length  int
}

func (helper *MedianHelper) justify() {
	middleLength := helper.Length / 2
	for helper.minHeap.Len() != middleLength {
		if helper.minHeap.Len() < middleLength {
			num := heap.Pop(helper.maxHeap)
			heap.Push(helper.minHeap, num)
		} else {
			num := heap.Pop(helper.minHeap)
			heap.Push(helper.maxHeap, num)
		}
	}
}

func (helper *MedianHelper) replace(src, target int) {
	if items, ok := helper.minHeap.ItemMaps[src]; ok && len(items) > 0 {
		heap.Remove(helper.minHeap, items[0].Index)
	} else {
		heap.Remove(helper.maxHeap, helper.maxHeap.ItemMaps[src][0].Index)
	}

	if helper.maxHeap.Len() == 0 || helper.maxHeap.Items[0].Value >= target {
		heap.Push(helper.maxHeap, target)
	} else {
		heap.Push(helper.minHeap, target)
	}

}

func (helper *MedianHelper) findMedianNum() float64 {
	helper.justify()
	if helper.Length%2 == 0 {
		return float64(helper.minHeap.Items[0].Value+helper.maxHeap.Items[0].Value) / 2.0
	}

	return float64(helper.maxHeap.Items[0].Value)
}

func NewMediaHelper(length int) *MedianHelper {
	return &MedianHelper{
		minHeap: NewMediaHeap(false),
		maxHeap: NewMediaHeap(true),
		Length:  length,
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	helper := NewMediaHelper(k)
	for i := 0; i < k; i++ {
		heap.Push(helper.maxHeap, nums[i])
	}

	ans := make([]float64, 0, len(nums)-k+1)
	ans = append(ans, helper.findMedianNum())

	for i := 0; i+k < len(nums); i++ {
		helper.replace(nums[i], nums[i+k])
		ans = append(ans, helper.findMedianNum())
	}

	return ans
}

// @lc code=end

func Test480() {
	// [1,-1,-1,3,5,6]。
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
