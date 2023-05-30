package section

import "fmt"

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	counter := make(map[byte][]int, 26)
	for i := 0; i < len(t); i++ {
		counter[t[i]] = append(counter[t[i]], i)
	}

	currentIndex := -1
	for i := 0; i < len(s); i++ {
		indexes, ok := counter[s[i]]
		if !ok {
			return false
		}

		if len(indexes) == 0 || indexes[len(indexes)-1] <= currentIndex {
			return false
		}

		for _, index := range indexes {
			if index > currentIndex {
				currentIndex = index
				break
			}
		}
	}

	return true
}

// @lc code=end

func Test392() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc") == false)
}
