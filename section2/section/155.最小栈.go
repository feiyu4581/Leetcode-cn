package section

import "fmt"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type StackNode struct {
	position int
	next     *StackNode
}

type MinStack struct {
	stacks []int
	length int
	size   int
	head   *StackNode
}

func Constructor() MinStack {
	return MinStack{
		stacks: make([]int, 12),
		length: 0,
		size:   12,
		head:   nil,
	}
}

func (this *MinStack) insort(position int) {
	if this.head == nil || this.stacks[this.head.position] >= this.stacks[position] {
		this.head = &StackNode{
			position: position,
			next:     this.head,
		}
	} else {
		current := this.head
		for {
			if current.next == nil || this.stacks[current.next.position] >= this.stacks[position] {
				current.next = &StackNode{
					position: position,
					next:     current.next,
				}
				break
			}
			current = current.next
		}
	}
}

func (this *MinStack) deleteSort(position int) {
	if this.head.position == position {
		this.head = this.head.next
	} else {
		current := this.head
		for {
			if current.next.position == position {
				current.next = current.next.next
				break
			}

			current = current.next
		}
	}
}

func (this *MinStack) Push(val int) {
	if this.length == this.size {
		stacks := make([]int, this.size*2)
		for i := 0; i < this.length; i++ {
			stacks[i] = this.stacks[i]
		}

		this.stacks = stacks
		this.size = this.size * 2
	}

	this.stacks[this.length] = val
	this.insort(this.length)

	this.length += 1
}

func (this *MinStack) Pop() {
	this.length -= 1
	this.deleteSort(this.length)
}

func (this *MinStack) Top() int {
	return this.stacks[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.stacks[this.head.position]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

func Test155() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin() == -3)
	minStack.Pop()
	fmt.Println(minStack.Top() == 0)
	fmt.Println(minStack.GetMin() == -2)
}
