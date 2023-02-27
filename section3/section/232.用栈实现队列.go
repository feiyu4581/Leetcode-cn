package section

import "fmt"

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	stack [100]int
	tail  int
	head  int
	len   int
}

func Constructor232() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack[this.tail] = x
	this.tail = (this.tail + 1) % 100
}

func (this *MyQueue) Pop() int {
	num := this.stack[this.head]
	this.head = (this.head + 1) % 100
	return num
}

func (this *MyQueue) Peek() int {
	return this.stack[this.head]
}

func (this *MyQueue) Empty() bool {
	return this.head == this.tail
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

func Test232() {
	myQueue := Constructor232()
	myQueue.Push(1)                       // queue is: [1]
	myQueue.Push(2)                       // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue.Peek())           // return 1
	fmt.Println(myQueue.Pop())            // return 1, queue is [2]
	fmt.Println(myQueue.Empty() == false) // return false
}
