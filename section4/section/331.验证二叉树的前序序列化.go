package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start

const (
	NodeType331 = iota
	LeafType331
)

func isValidSerialization(preorder string) bool {
	if preorder == "" {
		return false
	}

	if preorder == "#" {
		return true
	}

	items := strings.Split(preorder, ",")
	stacks := make([]int, 0)
	for i, item := range items {
		if item == "#" {
			stacks = append(stacks, LeafType331)
			for len(stacks) >= 3 {
				if stacks[len(stacks)-2] != LeafType331 || stacks[len(stacks)-3] != NodeType331 {
					break
				}
				stacks = stacks[:len(stacks)-3]
				if len(stacks) == 0 {
					return i == len(items)-1
				}
				stacks = append(stacks, LeafType331)
			}

		} else {
			stacks = append(stacks, NodeType331)
		}
	}

	return len(stacks) == 0
}

// @lc code=end

func Test331() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,#,2,#,6,#,#") == false)
	fmt.Println(isValidSerialization("#,#,#") == false)
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#") == false)
	fmt.Println(isValidSerialization("9,#,#,1") == false)
}
