package section

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution384 struct {
	current  []int
	original []int
}

func Constructor384(nums []int) Solution384 {
	current := make([]int, len(nums))
	for i := range nums {
		current[i] = nums[i]
	}
	return Solution384{
		current:  current,
		original: nums,
	}
}

func (this *Solution384) Reset() []int {
	for i := range this.original {
		this.current[i] = this.original[i]
	}

	return this.current
}

func (this *Solution384) Shuffle() []int {
	for i := len(this.current) - 1; i > 0; i-- {
		index := rand.Intn(i + 1)
		this.current[index], this.current[i] = this.current[i], this.current[index]
	}

	return this.current
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

func Test384() {
	solution := Constructor384([]int{1, 2, 3})
	fmt.Println(solution.Shuffle()) // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
	fmt.Println(solution.Reset())   // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
	fmt.Println(solution.Shuffle()) // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
}
