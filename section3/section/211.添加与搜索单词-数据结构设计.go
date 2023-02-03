package section

import "fmt"

/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start

type WordDictionary struct {
	Word  byte
	End   bool
	Child map[byte]*WordDictionary
}

func Constructor2() WordDictionary {
	return WordDictionary{
		Child: make(map[byte]*WordDictionary, 0),
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) > 0 {
		if nextTrie, ok := this.Child[word[0]]; !ok {
			this.Child[word[0]] = &WordDictionary{
				Word:  word[0],
				End:   len(word) == 1,
				Child: make(map[byte]*WordDictionary, 0),
			}
		} else if len(word) == 1 {
			nextTrie.End = true
		}

		this.Child[word[0]].AddWord(word[1:])
	}
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) > 0 {
		if word[0] == '.' {
			for _, nextTrie := range this.Child {
				if len(word) == 1 && nextTrie.End {
					return true
				}
				if nextTrie.Search(word[1:]) {
					return true
				}
			}
		} else {
			if nextTrie, ok := this.Child[word[0]]; ok {
				if len(word) == 1 && nextTrie.End {
					return true
				}

				return nextTrie.Search(word[1:])
			}
		}
	}

	return false
}

func Test211() {
	wordDictionary := Constructor2()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad") == false) // 返回 False
	fmt.Println(wordDictionary.Search("bad"))          // 返回 True
	fmt.Println(wordDictionary.Search(".ad"))          // 返回 True
	fmt.Println(wordDictionary.Search("b.."))          // 返回 True
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
