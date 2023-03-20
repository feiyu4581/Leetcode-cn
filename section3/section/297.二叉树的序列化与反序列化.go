package section

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
}

func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := make([]string, 0)

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0, len(nodes)*2)
		for _, node := range nodes {
			if node == nil {
				res = append(res, "nil")
			} else {
				res = append(res, fmt.Sprintf("%d", node.Val))
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}

		if len(newNodes) == 0 {
			res = res[:len(res)-len(nodes)]
		}
		nodes = newNodes
	}

	return strings.Join(res, ",")
}

func (this *Codec) generateNode(numStr string) *TreeNode {
	num, _ := strconv.Atoi(numStr)
	return &TreeNode{
		Val: num,
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	nums := strings.Split(data, ",")
	root := this.generateNode(nums[0])
	nodes := []*TreeNode{root}

	index := 1
	for index < len(nums) {
		newNode := make([]*TreeNode, 0)
		for _, node := range nodes {
			if nums[index] != "nil" {
				node.Left = this.generateNode(nums[index])
				newNode = append(newNode, node.Left)
			}

			index += 1

			if nums[index] != "nil" {
				node.Right = this.generateNode(nums[index])
				newNode = append(newNode, node.Right)
			}

			index += 1
		}

		nodes = newNode
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

func Test297() {
	ser := Constructor297()
	root := ToTreeNode([]int{1, 2, 3, 0, 0, 4, 5})
	data := ser.serialize(root)
	fmt.Println(data)
	fmt.Println(ser.deserialize(data).ToValues())
}
