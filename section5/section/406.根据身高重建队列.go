package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] && people[i][1] < people[j][1] {
			return true
		}

		return false
	})

	res := make([][]int, 0, len(people))
	for i := 0; i < len(people); i++ {
		gtCounts := people[i][1]
		index := 0
		for index < len(res) && gtCounts > 0 {
			if res[index][0] >= people[i][0] {
				gtCounts -= 1
			}

			index += 1
		}

		res = append(res[:index], append([][]int{people[i]}, res[index:]...)...)
	}

	return res
}

// @lc code=end

func Test406() {
	// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	// [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
	fmt.Println(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}))
}
