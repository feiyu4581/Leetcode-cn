package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

// @lc code=start

type Segment327 struct {
	Left  int
	Right int
	Val   int
}

func (segment *Segment327) IsInclude(position int) bool {
	return segment.Left <= position && position <= segment.Right
}

func (segment *Segment327) IsIncluded(left, right int) bool {
	return segment.Left >= left && right >= segment.Right
}

func (segment *Segment327) IsNoInclude(left, right int) bool {
	return segment.Left > right || segment.Right < left
}

func (segment *Segment327) IsEnd() bool {
	return segment.Left == segment.Right
}

func (segment *Segment327) Inc() {
	segment.Val += 1
}

type SegmentTree327 struct {
	Segments []Segment327
}

func NewSegmentTree327(left, right int) *SegmentTree327 {
	segmentTree327 := SegmentTree327{
		Segments: make([]Segment327, (right-left+1)*4),
	}

	segmentTree327.Build(1, left, right)
	return &segmentTree327
}

func (seg *SegmentTree327) Build(index, left, right int) {
	seg.Segments[index] = Segment327{
		Left:  left,
		Right: right,
	}

	if left == right {
		return
	}

	middle := left + (right-left)>>1
	seg.Build(index<<1, left, middle)
	seg.Build(index<<1|1, middle+1, right)

}

func (seg *SegmentTree327) Inc(index, position int) {
	if !seg.Segments[index].IsInclude(position) {
		return
	}

	seg.Segments[index].Val += 1
	if seg.Segments[index].IsEnd() {
		return
	}

	seg.Inc(index<<1, position)
	seg.Inc(index<<1|1, position)
}

func (seg *SegmentTree327) Query(index, left, right int) int {
	if seg.Segments[index].IsIncluded(left, right) {
		return seg.Segments[index].Val
	}

	if seg.Segments[index].IsEnd() || seg.Segments[index].IsNoInclude(left, right) {
		return 0
	}

	return seg.Query(index<<1, left, right) + seg.Query(index<<1|1, left, right)
}

func countRangeSum(nums []int, lower int, upper int) int {
	allNums := make([]int, 0, len(nums)*3+1)
	allNums = append(allNums, 0)

	prefixNums := make([]int, len(nums))
	for i := range nums {
		prefix := 0
		if i > 0 {
			prefix = prefixNums[i-1]
		}
		prefixNums[i] = nums[i] + prefix
		allNums = append(allNums, prefixNums[i]-upper, prefixNums[i]-lower, prefixNums[i])
	}

	numMaps := make(map[int]int, len(allNums))
	sort.Ints(allNums)
	count := 0
	for i := range allNums {
		if i == 0 || allNums[i] != allNums[i-1] {
			numMaps[allNums[i]] = count
			count += 1
		}
	}

	seg := NewSegmentTree327(0, count-1)
	seg.Inc(1, numMaps[0])
	res := 0
	for _, num := range prefixNums {
		res += seg.Query(1, numMaps[num-upper], numMaps[num-lower])
		seg.Inc(1, numMaps[num])
	}

	return res
}

// @lc code=end

func Test327() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2) == 3)
	fmt.Println(countRangeSum([]int{0}, 0, 0) == 1)
}
