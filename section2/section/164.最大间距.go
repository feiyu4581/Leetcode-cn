package section

import "fmt"

/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	// 稳定排序，因此对每个位数排序后一定最终有序
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			// 统计每个位置的个数
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			// 记录每个数字的最高位置
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			// 根据定位依次赋值
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func Test164() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}) == 3)
	fmt.Println(maximumGap([]int{10}) == 0)
}

// @lc code=end
