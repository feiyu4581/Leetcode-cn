package section

import "fmt"

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

type Trie struct {
	Word  byte
	End   bool
	Child map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		Child: make(map[byte]*Trie, 0),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) > 0 {
		if nextTrie, ok := this.Child[word[0]]; !ok {
			this.Child[word[0]] = &Trie{
				Word:  word[0],
				End:   len(word) == 1,
				Child: make(map[byte]*Trie, 0),
			}
		} else if len(word) == 1 {
			nextTrie.End = true
		}

		this.Child[word[0]].Insert(word[1:])
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) > 0 {
		if nextTrie, ok := this.Child[word[0]]; ok {
			if len(word) == 1 && nextTrie.End {
				return true
			}

			return nextTrie.Search(word[1:])
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) > 0 {
		if nextTrie, ok := this.Child[prefix[0]]; ok {
			if len(prefix) == 1 {
				return true
			}

			return nextTrie.StartsWith(prefix[1:])
		} else {
			return false
		}
	}

	return true
}

func Test208() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))        // 返回 True
	fmt.Println(trie.Search("app") == false) // 返回 False
	fmt.Println(trie.StartsWith("app"))      // 返回 True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 True
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix)
 */
// @lc code=end
