package section

import "fmt"

/*
 * @lc app=leetcode.cn id=352 lang=golang
 *
 * [352] 将数据流变为多个不相交区间
 */

// @lc code=start

type item352 struct {
	start int
	end   int
}

type SummaryRanges struct {
	ranges []*item352
}

func Constructor352() SummaryRanges {
	return SummaryRanges{ranges: []*item352{}}
}

func (this *SummaryRanges) findIndex(value int) (int, bool) {
	left, right := 0, len(this.ranges)-1
	for left <= right {
		midIndex := left + (right-left)>>1
		mid := this.ranges[midIndex]
		if mid.start <= value && mid.end >= value {
			return midIndex, true
		}

		if mid.start > value {
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}

	return left, false
}

func (this *SummaryRanges) AddNum(value int) {
	if index, ok := this.findIndex(value); !ok {
		newRanges := make([]*item352, 0, len(this.ranges)+1)
		if index > 0 {
			newRanges = append(newRanges, this.ranges[:index-1]...)
		}

		start, end := value, value
		if index > 0 {
			if this.ranges[index-1].end+1 == start {
				start = this.ranges[index-1].start
			} else {
				newRanges = append(newRanges, this.ranges[index-1])
			}
		}

		if index < len(this.ranges) {
			if this.ranges[index].start-1 == end {
				end = this.ranges[index].end
				index += 1
			}
		}

		newRanges = append(newRanges, &item352{start: start, end: end})
		if index < len(this.ranges) {
			newRanges = append(newRanges, this.ranges[index:]...)
		}
		this.ranges = newRanges
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, 0, len(this.ranges))
	for _, item := range this.ranges {
		res = append(res, []int{item.start, item.end})
	}

	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
// @lc code=end

func Test352() {
	summaryRanges := Constructor352()
	summaryRanges.AddNum(1) // arr = [1]
	// [[1, 1]]
	fmt.Println(summaryRanges.GetIntervals()) // 返回 [[1, 1]]
	summaryRanges.AddNum(3)                   // arr = [1, 3]
	// [[1, 1], [3, 3]]
	fmt.Println(summaryRanges.GetIntervals()) // 返回 [[1, 1], [3, 3]]
	summaryRanges.AddNum(7)                   // arr = [1, 3, 7]
	// [[1, 1], [3, 3], [7, 7]]
	fmt.Println(summaryRanges.GetIntervals()) // 返回 [[1, 1], [3, 3], [7, 7]]
	summaryRanges.AddNum(2)                   // arr = [1, 2, 3, 7]
	// [[1, 3], [7, 7]]
	fmt.Println(summaryRanges.GetIntervals()) // 返回 [[1, 3], [7, 7]]
	summaryRanges.AddNum(6)                   // arr = [1, 2, 3, 6, 7]
	// [[1, 3], [6, 7]]
	fmt.Println(summaryRanges.GetIntervals()) // 返回 [[1, 3], [6, 7]]
}
