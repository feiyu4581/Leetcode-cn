package section

import "fmt"

/*
 * @lc app=leetcode.cn id=336 lang=golang
 *
 * [336] 回文对
 */

// @lc code=start

type TrieNode336 struct {
	child  map[byte]*TrieNode336
	words  []int
	suffix []int
}

func NewTrieNode336() *TrieNode336 {
	return &TrieNode336{
		child: make(map[byte]*TrieNode336, 0),
	}
}

func (node *TrieNode336) AddWord(word byte) *TrieNode336 {
	if _, ok := node.child[word]; !ok {
		node.child[word] = NewTrieNode336()
	}

	return node.child[word]
}

func (node *TrieNode336) MatchWord(word byte) *TrieNode336 {
	return node.child[word]
}

func isPalindrome336(word string) bool {
	start, end := 0, len(word)-1
	for start < end {
		if word[start] != word[end] {
			return false
		}

		start += 1
		end -= 1
	}

	return true
}

func reverseString(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

func palindromePairs(words []string) [][]int {
	if len(words) == 0 {
		return nil
	}

	root := NewTrieNode336()
	for i := range words {
		cur := root
		reverseWord := reverseString(words[i])

		for j := 0; j < len(reverseWord); j++ {
			if isPalindrome336(reverseWord[j:]) {
				cur.suffix = append(cur.suffix, i)
			}
			cur = cur.AddWord(reverseWord[j])
		}

		cur.words = append(cur.words, i)
		cur.suffix = append(cur.suffix, i)
	}

	res := make([][]int, 0)
	for i := range words {
		cur := root
		start := 0
		for start < len(words[i]) && cur != nil {
			if len(cur.words) > 0 && isPalindrome336(words[i][start:]) {
				for _, matchWord := range cur.words {
					res = append(res, []int{i, matchWord})
				}
			}

			cur = cur.MatchWord(words[i][start])
			start += 1
		}

		if start == len(words[i]) && cur != nil {
			for _, matchWord := range cur.suffix {
				if matchWord != i {
					res = append(res, []int{i, matchWord})
				}
			}
		}
	}

	return res
}

// @lc code=end

func Test336() {
	//[[0,1],[1,0],[3,2],[2,4]]
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))

	// [[0,1],[1,0]]
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))

	// [[0,1],[1,0]]
	fmt.Println(palindromePairs([]string{"a", ""}))
}
