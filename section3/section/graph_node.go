package section

import (
	"math"
	"strconv"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func ToGraphNode(values [][]int) *Node {
	if len(values) == 0 {
		return nil
	}

	nodeMaps := make(map[int]*Node, len(values))
	for index := range values {
		nodeMaps[index+1] = &Node{Val: index + 1}
	}

	for index, neighbors := range values {
		for _, neighbor := range neighbors {
			nodeMaps[index+1].Neighbors = append(nodeMaps[index+1].Neighbors, nodeMaps[neighbor])
		}
	}

	return nodeMaps[1]
}

func (node *Node) String() string {
	return strconv.Itoa(node.Val)
}

func (node *Node) Values() [][]int {
	if node == nil {
		return nil
	}

	nodeList := make([]*Node, 100)
	nodes := []*Node{node}
	visisted := make(map[int]struct{}, 0)
	maxIndex := math.MinInt
	for len(nodes) > 0 {
		nextNodes := make([]*Node, 0)
		for _, node := range nodes {
			nodeList[node.Val] = node
			visisted[node.Val] = struct{}{}
			if node.Val > maxIndex {
				maxIndex = node.Val
			}
			for _, neighbor := range node.Neighbors {
				if _, ok := visisted[neighbor.Val]; ok {
					continue
				}

				nextNodes = append(nextNodes, neighbor)
			}
		}
		nodes = nextNodes
	}

	values := make([][]int, 0)
	for index := range nodeList {
		if index > maxIndex {
			break
		}

		node := nodeList[index]
		if node == nil {
			continue
		}

		neighborValues := make([]int, 0, len(node.Neighbors))
		for _, neighbor := range node.Neighbors {
			neighborValues = append(neighborValues, neighbor.Val)
		}

		values = append(values, neighborValues)
	}

	return values
}
