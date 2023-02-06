package section

import "fmt"

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start

type Trie212 struct {
	Children map[byte]*Trie212
	Word     string
}

var dfsPosition = [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func (trie *Trie212) Insert(word string) {
	current := trie
	for index := range word {
		if _, ok := current.Children[word[index]]; !ok {
			current.Children[word[index]] = &Trie212{
				Children: make(map[byte]*Trie212, 0),
			}
		}

		current = current.Children[word[index]]
	}

	current.Word = word
}

func findWords(board [][]byte, words []string) []string {
	trieRoot := &Trie212{Children: make(map[byte]*Trie212, 0)}
	for _, word := range words {
		trieRoot.Insert(word)
	}

	seen := make(map[string]struct{}, 0)
	var dfs func(node *Trie212, x, y int)
	dfs = func(trie *Trie212, x, y int) {
		ch := board[x][y]
		child, ok := trie.Children[board[x][y]]
		if !ok {
			return
		}

		if child.Word != "" {
			seen[child.Word] = struct{}{}
		}

		board[x][y] = '#'
		for _, position := range dfsPosition {
			nextX, nextY := x+position[0], y+position[1]
			if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) && board[nextX][nextY] != '#' {
				dfs(child, nextX, nextY)
			}
		}
		board[x][y] = ch
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(trieRoot, i, j)
		}
	}

	currents := make([]string, 0, len(seen))
	for word := range seen {
		currents = append(currents, word)
	}

	return currents
}

func Test212() {
	// ["eat", "oath"]
	fmt.Println(findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))

	// []
	fmt.Println(findWords([][]byte{
		{'a', 'a'},
	}, []string{"aaa"}))

	// []
	fmt.Println(findWords([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, []string{"abcd"}))
}

// @lc code=end
