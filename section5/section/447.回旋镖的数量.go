package section

import "fmt"

/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	ans := 0

	for _, p := range points {
		cnt := make(map[int]int, 0)
		for _, q := range points {
			diff := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[diff] += 1
		}

		for _, v := range cnt {
			ans += v * (v - 1)
		}
	}

	return ans
}

// @lc code=end

func Test447() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}) == 2)
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}) == 2)
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}}) == 0)
}
