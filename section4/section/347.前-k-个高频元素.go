package section

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

type item347 struct {
	value int
	count int
	index int
}

type frequent347 []*item347

func (f *frequent347) Push(x any) {
	item := x.(*item347)
	item.index = f.Len()
	*f = append(*f, item)
}

func (f *frequent347) Pop() any {
	temp := *f
	res := temp[len(temp)-1]
	temp = temp[:len(temp)-1]
	*f = temp

	return res
}

func (f frequent347) Len() int {
	return len(f)
}

func (f frequent347) Less(i, j int) bool {
	return f[i].count > f[j].count
}

func (f frequent347) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
	f[i].index = i
	f[j].index = j
}

func topKFrequent(nums []int, k int) []int {
	x := frequent347{}
	frequentMap := make(map[int]*item347)
	for _, num := range nums {
		if _, ok := frequentMap[num]; !ok {
			frequentMap[num] = &item347{
				value: num,
				count: 1,
			}
			heap.Push(&x, frequentMap[num])
		} else {
			frequentMap[num].count += 1
			heap.Fix(&x, frequentMap[num].index)
		}
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&x).(*item347).value)
	}

	return res
}

// @lc code=end

func Test347() {
	//  [-1,2]
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))

	//  [1,2]
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	//  [1]
	fmt.Println(topKFrequent([]int{1}, 1))
}
