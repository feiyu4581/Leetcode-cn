package section

import "fmt"

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type frontNode struct {
	node  *TreeNode
	depth int
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == root || q == root {
		return root
	}

	frontMaps := make(map[*TreeNode]*frontNode, 0)

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		fNode := &frontNode{
			node:  node,
			depth: depth,
		}
		if node.Left != nil {
			frontMaps[node.Left] = fNode
		}

		if node.Right != nil {
			frontMaps[node.Right] = fNode
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 1)
	left := &frontNode{
		node:  p,
		depth: frontMaps[p].depth + 1,
	}

	right := &frontNode{
		node:  q,
		depth: frontMaps[q].depth + 1,
	}

	for {
		if left.depth == right.depth && left.node == right.node {
			return left.node
		}

		if left.depth == right.depth {
			left, right = frontMaps[left.node], frontMaps[right.node]
		} else if left.depth > right.depth {
			left = frontMaps[left.node]
		} else {
			right = frontMaps[right.node]
		}
	}
}

// @lc code=end

func Test235() {
	root1 := ToTreeNode([]int{6, 1, 8, 1, 4, 7, 9, 0, 0, 3, 5})
	fmt.Println(lowestCommonAncestor(root1, root1.Left.Right, root1.Left.Right.Left).Val == 4)

	leftNode := ToTreeNode([]int{2, 0, 4})
	rightNode := ToTreeNode([]int{8, 7, 9})

	root := &TreeNode{
		Val:   6,
		Left:  leftNode,
		Right: rightNode,
	}

	fmt.Println(lowestCommonAncestor(root, leftNode, rightNode).Val == 6)

	leftNode = ToTreeNode([]int{2})
	rightNode = ToTreeNode([]int{1})

	leftNode.Left = rightNode
	fmt.Println(lowestCommonAncestor(leftNode, leftNode, rightNode).Val == 2)
}
