package section

import "fmt"

/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 */

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	s1Cnt, s2Cnt, index := 0, 0, 0
	recallMaps := make(map[int][2]int)
	for {
		s1Cnt += 1
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[index] {
				index += 1
				if index == len(s2) {
					s2Cnt += 1
					index = 0
				}
			}
		}

		if s1Cnt >= n1 {
			return s2Cnt / n2
		}

		if pre, ok := recallMaps[index]; ok {
			prefixS1Cnt, prefixS2Cnt := pre[0], pre[1]
			ans := pre[1] + (n1-prefixS1Cnt)/(s1Cnt-prefixS1Cnt)*(s2Cnt-prefixS2Cnt)
			nums := (n1 - prefixS1Cnt) % (s1Cnt - prefixS1Cnt)
			for i := 0; i < nums; i++ {
				for j := 0; j < len(s1); j++ {
					if s1[j] == s2[index] {
						index += 1
						if index == len(s2) {
							ans += 1
							index = 0
						}
					}
				}
			}

			return ans / n2
		} else {
			recallMaps[index] = [2]int{s1Cnt, s2Cnt}
		}
	}
}

// @lc code=end

func Test466() {
	fmt.Println(getMaxRepetitions("aaa", 3, "aa", 1) == 4)
	fmt.Println(getMaxRepetitions("acb", 4, "ab", 2) == 2)
	fmt.Println(getMaxRepetitions("acb", 1, "acb", 1) == 1)
}
