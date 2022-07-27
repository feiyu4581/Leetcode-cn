package section

import "fmt"

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	nums := make([]int, len(ratings))

	for index := range ratings {
		nums[index] = 1
	}

	for index := 1; index < len(ratings); index++ {
		if ratings[index] > ratings[index-1] {
			nums[index] = nums[index-1] + 1
		} else if ratings[index] < ratings[index-1] {
			for j := index; j >= 1; j-- {
				if ratings[j] < ratings[j-1] && nums[j] >= nums[j-1] {
					nums[j-1] = nums[j] + 1
				} else {
					break
				}
			}
		}
	}

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func Test135() {
	fmt.Println(candy([]int{1, 0, 2}) == 5)
	fmt.Println(candy([]int{1, 2, 2}) == 4)
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}) == 13)
}

// @lc code=end
