package section

import (
	"bytes"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

// @lc code=start

type Path388 struct {
	path        []byte
	IsFile      bool
	currentSize int
	childs      []*Path388
}

func (path *Path388) Add(name []byte, ladder int) int {
	child := &Path388{
		path:   name,
		childs: make([]*Path388, 0),
	}

	if bytes.ContainsRune(name, '.') {
		child.IsFile = true
	}

	parent := path
	for i := 0; i < ladder; i++ {
		parent = parent.childs[len(parent.childs)-1]
	}

	child.currentSize = parent.currentSize + 1 + len(name)
	parent.childs = append(parent.childs, child)

	if child.IsFile {
		return child.currentSize
	}

	return 0
}

func lengthLongestPath(input string) int {
	root := &Path388{childs: make([]*Path388, 0), currentSize: -1}

	index := 0
	nextToken := func() []byte {
		current := make([]byte, 0)
		if input[index] == 10 || input[index] == 9 {
			current = append(current, input[index])
			index += 1
			return current
		}

		for index < len(input) && input[index] != 9 && input[index] != 10 {
			current = append(current, input[index])
			index += 1
		}

		return current
	}

	ladder := 0
	maxSize := 0
	for index < len(input) {
		token := nextToken()
		switch token[0] {
		case 10:
			ladder = 0
		case 9:
			ladder += 1
		default:
			if size := root.Add(token, ladder); size > maxSize {
				maxSize = size
			}
		}
	}

	return maxSize
}

// @lc code=end

func Test388() {
	fmt.Println(lengthLongestPath(`dir
	subdir1
	subdir2
		file.ext`) == 20)
}
