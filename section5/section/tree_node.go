package section

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTreeNode(value int) *TreeNode {
	if value == 0 {
		return nil
	}

	return &TreeNode{Val: value}
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("[Tree]%d", node.Val)
}

func ToTreeNode(nodeValues []int) *TreeNode {
	if len(nodeValues) == 0 {
		return nil
	}

	head := generateTreeNode(nodeValues[0])
	currents := []*TreeNode{head}
	currentIndex := 0
	left := true
	for _, nodeValue := range nodeValues[1:] {
		if currentIndex >= len(currents) {
			nextCurrents := make([]*TreeNode, 0, len(currents)*2)
			for _, current := range currents {
				if current.Left != nil {
					nextCurrents = append(nextCurrents, current.Left)
				}

				if current.Right != nil {
					nextCurrents = append(nextCurrents, current.Right)
				}
			}

			currents = nextCurrents
			currentIndex = 0
		}

		if left {
			currents[currentIndex].Left = generateTreeNode(nodeValue)
			left = false
		} else {
			left = true
			currents[currentIndex].Right = generateTreeNode(nodeValue)
			currentIndex += 1
		}
	}

	return head
}

func (node *TreeNode) ToValues() []int {
	if node == nil {
		return nil
	}
	values := make([]int, 0)
	currents := []*TreeNode{node}
	currentIndex := 0

	for len(currents) > 0 {
		if currentIndex >= len(currents) {
			nextCurrents := make([]*TreeNode, 0, len(currents)*2)
			for _, current := range currents {
				if current.Left != nil {
					nextCurrents = append(nextCurrents, current.Left)
				}

				if current.Right != nil {
					nextCurrents = append(nextCurrents, current.Right)
				}
			}

			currents = nextCurrents
			currentIndex = 0
		}

		if currentIndex < len(currents) {
			values = append(values, currents[currentIndex].Val)
		}

		currentIndex += 1
	}
	return values
}
