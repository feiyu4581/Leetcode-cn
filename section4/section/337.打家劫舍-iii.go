package section

import "fmt"

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recursiveRob(root *TreeNode, includeSelf bool, cache map[*TreeNode][2]int) int {
	if root == nil {
		return 0
	}

	position := 0
	if !includeSelf {
		position = 1
	}
	if cacheRes, ok := cache[root]; ok {
		if cacheRes[position] != 0 {
			return cacheRes[position]
		}
	} else {
		cache[root] = [2]int{}
	}

	res := 0
	if includeSelf {
		res = root.Val + recursiveRob(root.Left, false, cache) + recursiveRob(root.Right, false, cache)
	}

	if num := recursiveRob(root.Left, true, cache) + recursiveRob(root.Right, true, cache); num > res {
		res = num
	}

	cacheRes := cache[root]
	cacheRes[position] = res
	cache[root] = cacheRes

	return res
}

func rob(root *TreeNode) int {
	return recursiveRob(root, true, make(map[*TreeNode][2]int))
}

// @lc code=end

func Test337() {
	fmt.Println(rob(ToTreeNode([]int{3, 2, 3, 0, 3, 0, 1})) == 7)
	fmt.Println(rob(ToTreeNode([]int{3, 4, 5, 1, 3, 0, 1})) == 9)
}
