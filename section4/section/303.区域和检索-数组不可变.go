package section

import "fmt"

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.nums[i]
	}

	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

func Test303() {
	numArray := Constructor303([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2) == 1)  // return 1 ((-2) + 0 + 3)
	fmt.Println(numArray.SumRange(2, 5) == -1) // return -1 (3 + (-5) + 2 + (-1))
	fmt.Println(numArray.SumRange(0, 5) == -3) // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
}
