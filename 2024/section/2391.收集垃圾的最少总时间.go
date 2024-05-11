package section

import "fmt"

/*
 * @lc app=leetcode.cn id=2391 lang=golang
 *
 * [2391] 收集垃圾的最少总时间
 */

// @lc code=start
func garbageCollection(garbage []string, travel []int) int {

	rangeGarbage := func(target rune) int {
		current := 0
		next := 0
		i := 0
		for i < len(garbage) {
			collectTime := 0
			for _, c := range garbage[i] {
				// 如果收集到目标垃圾，那么增加收集时间
				if c == target {
					collectTime++
				}
			}

			next += collectTime
			// 如果当前节点需要收集垃圾，那么加上迁移到当前节点的时间
			if collectTime > 0 {
				current += next
				next = 0
			}

			i += 1
			// next 记录迁移的时间，如果找到下一个需要收集的垃圾，那么就需要加上迁移时间
			if i < len(garbage) {
				next += travel[i-1]
			}
		}

		return current
	}

	return rangeGarbage('G') + rangeGarbage('P') + rangeGarbage('M')
}

// @lc code=end

func Test2391() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}) == 21)
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}) == 37)
}
