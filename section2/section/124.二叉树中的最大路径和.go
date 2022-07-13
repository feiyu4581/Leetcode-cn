package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

func maxPath(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func computeParentMaxPathSum(childNode *TreeNode, node *TreeNode, parentMap map[*TreeNode]*TreeNode,
	cache map[*TreeNode]int, parentLeftCache map[*TreeNode]int, parentRightCache map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	rootNode := node.Right
	parentCache := parentLeftCache
	if childNode == node.Right {
		rootNode = node.Left
		parentCache = parentRightCache
	}

	if pathSum, ok := parentCache[node]; ok {
		return pathSum
	}

	pathSum := maxPath(
		computeParentMaxPathSum(node, parentMap[node], parentMap, cache, parentLeftCache, parentRightCache),
		computeRootMaxPathSum(rootNode, cache),
	) + node.Val

	parentCache[node] = pathSum
	return pathSum
}

func computeRootMaxPathSum(node *TreeNode, cache map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	if pathSum, ok := cache[node]; ok {
		return pathSum
	}

	pathSum := maxPath(
		computeRootMaxPathSum(node.Left, cache),
		computeRootMaxPathSum(node.Right, cache),
	)

	pathSum = maxPath(node.Val, node.Val+pathSum)
	cache[node] = pathSum
	return pathSum
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	parentMap := make(map[*TreeNode]*TreeNode, 0)
	nodes := []*TreeNode{root}

	start, length := 0, len(nodes)
	for {
		offset := 0
		for i := start; i < length; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
				parentMap[nodes[i].Left] = nodes[i]
				offset += 1
			}

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
				parentMap[nodes[i].Right] = nodes[i]
				offset += 1
			}
		}

		if offset == 0 {
			break
		}

		start = length
		length = length + offset
	}

	rootCache := make(map[*TreeNode]int, 0)
	parentLeftCache := make(map[*TreeNode]int, 0)
	parentRightCache := make(map[*TreeNode]int, 0)
	maxValue := math.MinInt
	for _, node := range nodes {
		value := maxPath(
			computeRootMaxPathSum(node, rootCache),
			computeParentMaxPathSum(node, parentMap[node], parentMap, rootCache, parentLeftCache, parentRightCache)+node.Val,
		)
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func Test124() {
	fmt.Println(maxPathSum(ToTreeNode([]int{-1, -2, 10, -6, 0, -3, -6})) == 10)
	fmt.Println(maxPathSum(ToTreeNode([]int{1, 2, 3})) == 6)
	fmt.Println(maxPathSum(ToTreeNode([]int{-10, 9, 20, 0, 0, 15, 7})) == 42)
	fmt.Println(maxPathSum(ToTreeNode([]int{1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0})) == 4)
}

// @lc code=end
