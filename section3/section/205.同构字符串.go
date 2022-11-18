package section

import "fmt"

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	charMap := make(map[uint8]uint8, 0)
	targetMap := make(map[uint8]struct{}, 0)
	if len(s) != len(t) {
		return false
	}

	for index := range s {
		if target, ok := charMap[s[index]]; ok {
			if target != t[index] {
				return false
			}
		} else {
			if _, ok := targetMap[t[index]]; ok {
				return false
			}

			charMap[s[index]] = t[index]
			targetMap[t[index]] = struct{}{}
		}
	}

	return true
}

func Test205() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar") == false)
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba") == false)
}

// @lc code=end
