package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start

func removeDuplicateLetters(s string) string {
	inUse := make(map[byte]bool, 0)
	stacks := make([]byte, 0)

	counter := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counter[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		counter[s[i]] -= 1
		if ok := inUse[s[i]]; ok {
			continue
		}

		for len(stacks) > 0 && stacks[len(stacks)-1] > s[i] {
			if counter[stacks[len(stacks)-1]] == 0 {
				break
			}

			inUse[stacks[len(stacks)-1]] = false
			stacks = stacks[:len(stacks)-1]
		}

		stacks = append(stacks, s[i])
		inUse[s[i]] = true
	}

	return string(stacks)
}

// @lc code=end

func Test316() {
	fmt.Println(removeDuplicateLetters("cdadabcc") == "adbc")
	fmt.Println(removeDuplicateLetters("bcabc") == "abc")
	fmt.Println(removeDuplicateLetters("cbacdcbc") == "acdb")
}
