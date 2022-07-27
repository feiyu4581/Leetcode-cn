package section

import "fmt"

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	index, current := -1, 0
	allDiff := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		// 存储以 i 为终点的时候，需要额外多消耗的油
		if i == 0 {
			allDiff[i] = diff
		} else {
			allDiff[i] = allDiff[i-1] + diff
		}

		// 如果起点未选择，并且当前节点启动的时候油有剩余，那么从当前节点开始
		if index == -1 && diff >= 0 {
			current = diff
			index = i
		} else if index >= 0 {
			// 如果节点已经选择，并且加上当前节点油不够，那么重新开始选择节点
			current = current + diff
			if current < 0 {
				index = -1
			}
		}
	}

	// 如果节点已经选择，并且起点到最后一个节点剩余的油满足从从零到起点位置的油，那么认为找到结果值
	if index == 0 && current >= 0 {
		return index
	} else if index > 0 && current+allDiff[index-1] >= 0 {
		return index
	}

	return -1
}

func Test134() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}) == 3)
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}) == -1)
}

// @lc code=end
