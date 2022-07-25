package section

import "fmt"

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a NextNode.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func recursiveCloneGraph(node *Node, visited map[*Node]*Node) *Node {
	if newNode, ok := visited[node]; ok {
		return newNode
	}

	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, recursiveCloneGraph(neighbor, visited))
	}

	return newNode
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	return recursiveCloneGraph(node, make(map[*Node]*Node, 0))
}

func Test133() {
	node := ToGraphNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	fmt.Println(cloneGraph(node).Values())
}

// @lc code=end
