package section

import "fmt"

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	counter := make(map[byte]int, len(ransomNote))
	for i := 0; i < len(magazine); i++ {
		counter[magazine[i]]++
	}

	for j := 0; j < len(ransomNote); j++ {
		if counter[ransomNote[j]] > 0 {
			counter[ransomNote[j]]--
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

func Test383() {
	fmt.Println(canConstruct("a", "b") == false)
	fmt.Println(canConstruct("aa", "ab") == false)
	fmt.Println(canConstruct("aa", "aab") == true)
}
