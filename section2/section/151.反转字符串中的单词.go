package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	words := make([]string, 0)
	currentWord := make([]int32, 0)
	for _, char := range s {
		if char != ' ' {
			currentWord = append(currentWord, char)
		} else if len(currentWord) > 0 {
			words = append(words, string(currentWord))
			currentWord = make([]int32, 0)
		}
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	if len(words) == 0 {
		return ""
	}

	results := strings.Builder{}
	results.WriteString(words[len(words)-1])
	for i := len(words) - 2; i >= 0; i-- {
		results.WriteRune(' ')
		results.WriteString(words[i])
	}

	return results.String()
}

func Test151() {
	fmt.Println(reverseWords("the sky is blue") == "blue is sky the")
	fmt.Println(reverseWords("  hello world  ") == "world hello")
	fmt.Println(reverseWords("a good   example") == "example good a")
}

// @lc code=end
