package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray307 struct {
	nums []int
	size int
	sums []int
}

func Constructor307(nums []int) NumArray307 {
	size := int(math.Sqrt(float64(len(nums))))

	sums := make([]int, len(nums)/size+1)
	for i := range nums {
		sums[i/size] += nums[i]
	}

	return NumArray307{
		nums: nums,
		size: size,
		sums: sums,
	}
}

func (this *NumArray307) Update(index int, val int) {
	this.sums[index/this.size] += val - this.nums[index]
	this.nums[index] = val

}

func (this *NumArray307) SumRange(left int, right int) int {
	leftSum, rightSum := left/this.size, right/this.size

	if leftSum == rightSum {
		sum := 0
		for i := left; i <= right; i++ {
			sum += this.nums[i]
		}

		return sum
	}

	res := this.SumRange(left, leftSum*this.size+this.size-1) + this.SumRange(rightSum*this.size, right)
	for i := leftSum + 1; i < rightSum; i++ {
		res += this.sums[i]
	}

	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end

func Test307() {
	numArray := Constructor307([]int{1, 3, 5, 7, 9})
	fmt.Println(numArray.SumRange(0, 2) == 9)
	numArray.Update(1, 2)
	fmt.Println(numArray.SumRange(0, 2) == 8)
}
