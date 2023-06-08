package section

import "fmt"

type NextNode struct {
	Val   int
	Next  *NextNode
	Left  *NextNode
	Right *NextNode
}

func generateNode(value int) *NextNode {
	if value == 0 {
		return nil
	}

	return &NextNode{Val: value}
}

func (node *NextNode) String() string {
	return fmt.Sprintf("[NextNode]%d", node.Val)
}

func ToNode(nodeValues []int) *NextNode {
	if len(nodeValues) == 0 {
		return nil
	}

	head := generateNode(nodeValues[0])
	currents := []*NextNode{head}
	currentIndex := 0
	left := true
	for _, nodeValue := range nodeValues[1:] {
		if currentIndex >= len(currents) {
			nextCurrents := make([]*NextNode, 0, len(currents)*2)
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
			currents[currentIndex].Left = generateNode(nodeValue)
			left = false
		} else {
			left = true
			currents[currentIndex].Right = generateNode(nodeValue)
			currentIndex += 1
		}
	}

	return head
}

func (node *NextNode) ToValues() []int {
	values := make([]int, 0)
	currents := []*NextNode{node}
	currentIndex := 0

	for len(currents) > 0 {
		if currentIndex >= len(currents) {
			nextCurrents := make([]*NextNode, 0, len(currents)*2)
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
