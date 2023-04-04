package section

import "fmt"

/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	counter := make([]map[byte]struct{}, len(words))
	repeatedMap := make(map[int]map[int]bool, len(words))
	for i := range words {
		counter[i] = make(map[byte]struct{}, 26)
		repeatedMap[i] = make(map[int]bool, len(words))

		for j := range words[i] {
			counter[i][words[i][j]] = struct{}{}
		}
	}

	for i := 0; i < 26; i++ {
		char := byte(i + 'a')
		nums := make([]int, 0)
		for index := range words {
			if _, ok := counter[index][char]; ok {
				nums = append(nums, index)
			}
		}

		for first := range nums {
			for last := first + 1; last < len(nums); last++ {
				repeatedMap[nums[first]][nums[last]] = true
			}
		}
	}

	maxValue := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if !repeatedMap[i][j] {
				if length := len(words[i]) * len(words[j]); length > maxValue {
					maxValue = length
				}
			}
		}
	}

	return maxValue
}

// @lc code=end

func Test318() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}) == 16)
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}) == 4)
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}) == 0)
}
