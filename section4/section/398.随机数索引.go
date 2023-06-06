package section

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
type Solution398 struct {
	counter map[int][]int
}

func Constructor398(nums []int) Solution398 {
	counter := make(map[int][]int)
	for index, num := range nums {
		counter[num] = append(counter[num], index)
	}

	return Solution398{counter: counter}
}

func (this *Solution398) Pick(target int) int {
	if nums, ok := this.counter[target]; ok {
		return nums[rand.Intn(len(nums))]
	}

	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

func Test398() {
	solution := Constructor398([]int{1, 2, 3, 3, 3})
	fmt.Println(solution.Pick(3)) // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
	fmt.Println(solution.Pick(1)) // 返回 0 。因为只有 nums[0] 等于 1 。
	fmt.Println(solution.Pick(3)) // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
}
