package section

import "fmt"

/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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

type BSTIterator struct {
	root    *TreeNode
	stacks  []*TreeNode
	current *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stacks := make([]*TreeNode, 0)
	current := root
	for current.Left != nil {
		stacks = append(stacks, current)
		current = current.Left
	}

	return BSTIterator{
		root:    root,
		stacks:  stacks,
		current: current,
	}
}

func (this *BSTIterator) Next() (val int) {
	val = this.current.Val
	if this.current.Right != nil {
		this.current = this.current.Right
		for this.current.Left != nil {
			this.stacks = append(this.stacks, this.current)
			this.current = this.current.Left
		}
	} else if len(this.stacks) > 0{
		this.current = this.stacks[len(this.stacks) - 1]
		this.stacks = this.stacks[:len(this.stacks) - 1]
	} else {
		this.current = nil
	}
	return
}

func (this *BSTIterator) HasNext() bool {
	return this.current != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

func Test173() {
	iterator := Constructor(ToTreeNode([]int{7, 3, 15, 0, 0, 9, 20}))
	fmt.Println(iterator.Next() == 3)
	fmt.Println(iterator.Next() == 7)
	fmt.Println(iterator.HasNext())

	fmt.Println(iterator.Next() == 9)
	fmt.Println(iterator.HasNext())

	fmt.Println(iterator.Next() == 15)
	fmt.Println(iterator.HasNext())

	fmt.Println(iterator.Next() == 20)
	fmt.Println(iterator.HasNext() == false)
}
