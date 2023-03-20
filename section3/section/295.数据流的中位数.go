package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start

type H295 struct {
	data []int
	min  bool
}

func (h *H295) Len() int {
	return len(h.data)
}

func (h *H295) Less(i, j int) bool {
	if h.min {
		return h.data[i] < h.data[j]
	}

	return h.data[i] > h.data[j]
}
func (h *H295) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *H295) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *H295) Pop() interface{} {
	num := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return num
}

type MedianFinder struct {
	minHeap *H295
	maxHeap *H295
}

func Constructor295() MedianFinder {
	minHeap := &H295{min: true}
	maxHeap := &H295{min: false}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}

func (this *MedianFinder) isBalance() bool {
	length := this.minHeap.Len() + this.maxHeap.Len()
	if length%2 == 0 {
		return this.minHeap.Len() == this.maxHeap.Len()
	}

	return this.minHeap.Len()+1 == this.maxHeap.Len()
}

func (this *MedianFinder) computeBalance() {
	for !this.isBalance() {
		if this.minHeap.Len() >= this.maxHeap.Len() {
			num := heap.Pop(this.minHeap)
			heap.Push(this.maxHeap, num)
		} else {
			num := heap.Pop(this.maxHeap)
			heap.Push(this.minHeap, num)
		}
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= this.maxHeap.data[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	this.computeBalance()
}

func (this *MedianFinder) FindMedian() float64 {
	length := this.minHeap.Len() + this.maxHeap.Len()
	if length%2 == 0 {
		return float64(this.maxHeap.data[0]+this.minHeap.data[0]) / 2
	}

	return float64(this.maxHeap.data[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

func Test295() {
	medianFinder := Constructor295()
	medianFinder.AddNum(1) // arr = [1]
	medianFinder.AddNum(2) // arr = [1, 2]
	// 1.5
	fmt.Println(medianFinder.FindMedian() == 1.5) // 返回 1.5 ((1 + 2) / 2)
	medianFinder.AddNum(3)                        // arr[1, 2, 3]
	// 2.0
	fmt.Println(medianFinder.FindMedian() == 2) // return 2.0
}
