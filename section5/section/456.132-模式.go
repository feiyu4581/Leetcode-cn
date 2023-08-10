package section

import "fmt"

/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132 模式
 */

// @lc code=start
func find132pattern(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}

	uphills := make([][2]int, 0)

	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if i > start+1 {
				left, right := nums[start], nums[i-1]
				uphillBorder := len(uphills)
				for i := len(uphills) - 1; i >= 0; i-- {
					if uphills[i][0] >= right || uphills[i][1] <= left {
						break
					}
					uphillBorder -= 1
				}

				uphills = uphills[:uphillBorder]
				uphills = append(uphills, [2]int{left, right})
			}

			start = i
		}

		for _, uphill := range uphills {
			if nums[i] > uphill[0] && nums[i] < uphill[1] {
				return true
			}
		}
	}

	return false
}

// @lc code=end

func Test456() {
	fmt.Println(find132pattern([]int{80, 84, 70, 80, 60, 70, 50, 60, 82, 85}) == true)
	fmt.Println(find132pattern([]int{0, -1000, 2000, -3000, 4000, -5000, 6000, -7000, 8000, -9000, 10000, -11000, 12000, -13000, 14000}) == false)
	fmt.Println(find132pattern([]int{1, 2, 3, 4}) == false)
	fmt.Println(find132pattern([]int{3, 1, 4, 2}) == true)
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}) == true)
}
