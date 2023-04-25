package section

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {
	return false
}
func (this NestedInteger) GetInteger() int {
	return 0
}
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	current         int
	currentIterator *NestedIterator
	nestedList      []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		current:    0,
		nestedList: nestedList,
	}
}

func (this *NestedIterator) Next() int {
	res := 0
	if this.currentIterator != nil {
		res = this.currentIterator.Next()
		if !this.currentIterator.HasNext() {
			this.current += 1
			this.currentIterator = nil
		}

		return res
	}

	if this.nestedList[this.current].IsInteger() {
		res = this.nestedList[this.current].GetInteger()
		this.current += 1
	} else {
		this.currentIterator = Constructor(this.nestedList[this.current].GetList())
		return this.Next()
	}

	return res
}

func (this *NestedIterator) HasNext() bool {
	if this.current >= len(this.nestedList) {
		return false
	}

	if this.nestedList[this.current].IsInteger() {
		return true
	}

	if this.currentIterator == nil {
		this.currentIterator = Constructor(this.nestedList[this.current].GetList())
	}

	if this.currentIterator.HasNext() {
		return true
	}

	this.currentIterator = nil
	this.current += 1
	return this.HasNext()
}

// @lc code=end

func Test341() {

}
