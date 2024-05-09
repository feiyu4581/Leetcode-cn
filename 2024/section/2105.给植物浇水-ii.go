package section

import "fmt"

/*
 * @lc app=leetcode.cn id=2105 lang=golang
 *
 * [2105] 给植物浇水 II
 */

// @lc code=start
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	count := 0
	left, leftCapacityA, right, rightCapacityB := 0, capacityA, len(plants)-1, capacityB
	for left <= right {
		if left == right {
			if leftCapacityA < plants[left] && rightCapacityB < plants[left] {
				count += 1
			}
			break
		}

		if leftCapacityA >= plants[left] {
			leftCapacityA -= plants[left]
		} else {
			leftCapacityA = capacityA - plants[left]
			count += 1
		}

		if rightCapacityB >= plants[right] {
			rightCapacityB -= plants[right]
		} else {
			rightCapacityB = capacityB - plants[right]
			count += 1
		}

		left += 1
		right -= 1
	}

	return count
}

// @lc code=end

func Test2105() {
	fmt.Println(minimumRefill([]int{2, 2, 3, 3}, 5, 5) == 1)
	fmt.Println(minimumRefill([]int{2, 2, 3, 3}, 3, 4) == 2)
	fmt.Println(minimumRefill([]int{5}, 10, 8) == 0)
	fmt.Println(minimumRefill([]int{2, 1, 1}, 2, 2) == 0)
}
