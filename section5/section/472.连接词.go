package section

import "fmt"

/*
 * @lc app=leetcode.cn id=472 lang=golang
 *
 * [472] 连接词
 */

// @lc code=start

type TrieWord struct {
	Child map[byte]*TrieWord
	Word  string
	Cache map[string]int
}

func NewTrieWord() *TrieWord {
	return &TrieWord{
		Child: make(map[byte]*TrieWord),
	}
}

func (trie *TrieWord) Add(word string) {
	current := trie
	for i := 0; i < len(word); i++ {
		if _, ok := current.Child[word[i]]; !ok {
			current.Child[word[i]] = NewTrieWord()
		}
		current = current.Child[word[i]]
	}

	current.Word = word
}

func (trie *TrieWord) FindConcatenated(word string) (num int) {
	if _, ok := trie.Cache[word]; ok {
		return trie.Cache[word]
	}

	defer func() {
		trie.Cache[word] = num
	}()

	current := trie
	index := 0
	for index < len(word) {
		if _, ok := current.Child[word[index]]; !ok {
			return 0
		}

		current = current.Child[word[index]]
		if current.Word != "" && index+1 < len(word) && trie.FindConcatenated(word[index+1:]) >= 1 {
			return 2
		}

		index += 1
	}

	if current.Word != "" {
		return 1
	}

	return 0
}

func findAllConcatenatedWordsInADict(words []string) []string {
	trie := NewTrieWord()
	trie.Cache = make(map[string]int, len(words))
	for _, word := range words {
		trie.Add(word)
	}

	ans := make([]string, 0)
	for _, word := range words {
		if trie.FindConcatenated(word) >= 2 {
			ans = append(ans, word)
		}
	}

	return ans
}

// @lc code=end

func Test472() {
	// ["catsdogcats","dogcatsdog","ratcatdogcat"]
	fmt.Println(findAllConcatenatedWordsInADict([]string{
		"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat",
	}))

	// ["catdog"]
	fmt.Println(findAllConcatenatedWordsInADict([]string{
		"cat", "dog", "catdog",
	}))
}
