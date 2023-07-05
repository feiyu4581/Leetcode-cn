package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start

type repeat424 struct {
	start int
	end   int
}

func (repeat *repeat424) Length() int {
	return repeat.end - repeat.start + 1
}

func characterReplacement(s string, k int) int {
	lastIndex := 0
	set := make(map[byte][]repeat424, 26)
	for i := 1; i < len(s); i++ {
		if s[i] != s[lastIndex] {
			set[s[lastIndex]] = append(set[s[lastIndex]], repeat424{
				start: lastIndex,
				end:   i - 1,
			})

			lastIndex = i
		}
	}

	set[s[lastIndex]] = append(set[s[lastIndex]], repeat424{
		start: lastIndex,
		end:   len(s) - 1,
	})

	res := 0
	for _, repeats := range set {
		for i := range repeats {
			num := repeats[i].Length()
			diff := k
			for j := i + 1; j < len(repeats); j++ {
				if distance := repeats[j].start - repeats[j-1].end - 1; distance <= diff {
					num += repeats[j].Length() + distance
					diff -= distance
				} else {
					break
				}
			}

			num = num + diff
			if num > len(s) {
				num = len(s)
			}

			if num > res {
				res = num
			}
		}
	}

	return res
}

// @lc code=end

func Test424() {
	fmt.Println(characterReplacement("ABBB", 2) == 4)
	fmt.Println(characterReplacement("ABAB", 2) == 4)
	fmt.Println(characterReplacement("AABABBA", 1) == 4)
}
