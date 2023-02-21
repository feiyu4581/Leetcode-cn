package section

import "fmt"

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	stack []int
}

func Constructor225() MyStack {
	return MyStack{stack: make([]int, 0, 10)}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	num := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return num
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

func Test225() {
	myStack := Constructor225()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top() == 2)       // 返回 2
	fmt.Println(myStack.Pop() == 2)       // 返回 2
	fmt.Println(myStack.Empty() == false) // 返回 False
	fmt.Println(myStack.Pop() == 1)       // 返回 2
	fmt.Println(myStack.Empty() == true)  // 返回 False
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
