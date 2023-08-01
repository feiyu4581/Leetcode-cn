package section

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
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

func Constructor449() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	values := make([]string, 0)
	nodes := []*TreeNode{root}

	values = append(values, fmt.Sprintf("%d", root.Val))

	for len(nodes) > 0 {
		nextNodes := make([]*TreeNode, 0)

		nilCount := 0
		for _, node := range nodes {
			if node.Left != nil {
				if nilCount > 0 {
					values = append(values, fmt.Sprintf("nil-%d", nilCount))
					nilCount = 0
				}

				nextNodes = append(nextNodes, node.Left)
				values = append(values, fmt.Sprintf("%d", node.Left.Val))
			} else {
				nilCount += 1
			}

			if node.Right != nil {
				if nilCount > 0 {
					values = append(values, fmt.Sprintf("nil-%d", nilCount))
					nilCount = 0
				}

				nextNodes = append(nextNodes, node.Right)
				values = append(values, fmt.Sprintf("%d", node.Right.Val))
			} else {
				nilCount += 1
			}
		}

		if nilCount > 0 {
			values = append(values, fmt.Sprintf("nil-%d", nilCount))
			nilCount = 0
		}
		nodes = nextNodes
	}

	return strings.Join(values, ",")
}

func (this *Codec) newNode(data string) *TreeNode {
	val, _ := strconv.Atoi(data)
	return &TreeNode{Val: val}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	nodes := strings.Split(data, ",")

	root := this.newNode(nodes[0])
	currents := []*TreeNode{root}

	index := 1
	for len(currents) > 0 || index < len(nodes) {
		nextCurrents := make([]*TreeNode, 0)

		nilCount := 0
		for _, node := range currents {
			if index >= len(nodes) {
				break
			}
			if nilCount == 0 && strings.HasPrefix(nodes[index], "nil") {
				nilCount, _ = strconv.Atoi(strings.Split(nodes[index], "-")[1])
				index += 1
			}

			if nilCount >= 2 {
				nilCount -= 2
				continue
			}

			if nilCount == 1 {
				node.Right = this.newNode(nodes[index])
				nextCurrents = append(nextCurrents, node.Right)
				index += 1
				nilCount = 0
			} else {
				node.Left = this.newNode(nodes[index])
				nextCurrents = append(nextCurrents, node.Left)
				index += 1

				if index >= len(nodes) {
					break
				}

				if nilCount == 0 && strings.HasPrefix(nodes[index], "nil") {
					nilCount, _ = strconv.Atoi(strings.Split(nodes[index], "-")[1])
					nilCount -= 1
					index += 1

					continue
				}

				node.Right = this.newNode(nodes[index])
				nextCurrents = append(nextCurrents, node.Right)
				index += 1
			}
		}

		currents = nextCurrents
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

func Test449() {
	ser := Constructor449()
	// [2, 1, 3]
	//fmt.Println(ser.deserialize(ser.serialize(ToTreeNode([]int{2, 1, 3, 0, 2}))).ToValues())

	// []
	//fmt.Println(ser.deserialize(ser.serialize(ToTreeNode([]int{}))).ToValues())
	// [2, 1]
	fmt.Println(ser.deserialize(ser.serialize(ToTreeNode([]int{2, 1}))).ToValues())
}
