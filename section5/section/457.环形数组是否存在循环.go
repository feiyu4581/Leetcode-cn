package section

import "fmt"

/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组是否存在循环
 */

// @lc code=start
func circularArrayLoop(nums []int) bool {
	visited := make([]int, len(nums))
	turn := 1
	for i := 0; i < len(nums); i++ {
		if visited[i] == 0 {
			flag, current := nums[i] > 0, i
			for {
				visited[current] = turn
				next := (nums[current] + current) % len(nums)
				if next < 0 {
					next = len(nums) + next
				}

				if current == next {
					break
				}

				if visited[next] == turn {
					return true
				}

				newFlag := nums[next] > 0
				if newFlag != flag || visited[next] != 0 {
					break
				}

				current = next
			}

			turn += 1
		}
	}

	return false
}

// @lc code=end

func Test457() {
	fmt.Println(circularArrayLoop([]int{-1, -2, -3, -4, -5}) == false)
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}) == true)
	fmt.Println(circularArrayLoop([]int{-1, 2}) == false)
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}) == false)
}
