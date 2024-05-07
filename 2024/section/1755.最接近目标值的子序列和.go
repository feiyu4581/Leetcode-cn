package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1755 lang=golang
 *
 * [1755] 最接近目标值的子序列和
 */

// @lc code=start

func dfsMinAbsDifference(nums []int, start, end, sum int, record []int) {
	record = append(record, sum)
	for i := start; i < end; i++ {
		dfsMinAbsDifference(nums, i+1, end, sum+nums[i], record)
	}
}

func minAbsDifference(nums []int, goal int) int {
	mid := len(nums) >> 1
	left := make([]int, 1<<mid)
	right := make([]int, 1<<(len(nums)-mid))

	index := 0
	var dfsMinAbsDifference func(start, end, sum int, record []int)
	dfsMinAbsDifference = func(start, end, sum int, record []int) {
		record[index] = sum
		index += 1
		for i := start; i < end; i++ {
			dfsMinAbsDifference(i+1, end, sum+nums[i], record)
		}
	}

	dfsMinAbsDifference(0, mid, 0, left)
	index = 0
	dfsMinAbsDifference(mid, len(nums), 0, right)

	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	target := abs(goal)
	for _, v := range left {
		if target = min(target, abs(goal-v)); target == 0 {
			return 0
		}
	}

	for _, v := range right {
		if target = min(target, abs(goal-v)); target == 0 {
			return 0
		}
	}

	sort.Ints(left)
	sort.Ints(right)
	l, r := 0, len(right)-1
	for l < len(left) && r >= 0 {
		sum := left[l] + right[r]
		target = min(target, abs(goal-sum))
		if target == 0 {
			return 0
		}

		if sum > goal {
			r -= 1
		} else {
			l += 1
		}
	}

	return target
}

// @lc code=end

func Test1755() {
	fmt.Println(minAbsDifference([]int{5, -7, 3, 5}, 6) == 0)
	fmt.Println(minAbsDifference([]int{7, -9, 15, -2}, -5) == 1)
	fmt.Println(minAbsDifference([]int{1, 2, 3}, -7) == 7)
	fmt.Println(minAbsDifference([]int{-7074297, 3076735, -5846354, 5008659, -126683, 7039557, 6708811, 3189666, -6102417, 6078975, -6448946, -4995910, 2964239, -3248847, -4392269, 7473223, -1356059, 3978911, 8009187, -316441, 6524770, 8280309, -2798383, 1310839, 6306594, -6548611, -9712711, 1639314}, 493409180) == 415819391)
}
