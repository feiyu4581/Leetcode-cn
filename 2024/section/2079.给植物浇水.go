package section

import "fmt"

/*
 * @lc app=leetcode.cn id=2079 lang=golang
 *
 * [2079] 给植物浇水
 */

// @lc code=start
func wateringPlants(plants []int, capacity int) int {
	count := 0

	currentCapacity := capacity
	position, index := 0, 1
	for index <= len(plants) {
		if currentCapacity >= plants[index-1] {
			currentCapacity -= plants[index-1]

			count += index - position
			position = index

			index += 1
		} else {
			count += index - 1
			currentCapacity = capacity
			position = 0
		}
	}

	return count
}

// @lc code=end

func Test2079() {
	fmt.Println(wateringPlants([]int{2, 2, 3, 3}, 5) == 14)
	fmt.Println(wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4) == 30)
	fmt.Println(wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8) == 49)
}
