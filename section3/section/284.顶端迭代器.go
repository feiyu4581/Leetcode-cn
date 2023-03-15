package section

/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [284] 顶端迭代器
 */

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return false
}

func (this *Iterator) next() int {
	return 0
}

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter     *Iterator
	currents []int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:     iter,
		currents: make([]int, 0),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.currents) > 0 || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if len(this.currents) > 0 {
		num := this.currents[len(this.currents)-1]
		this.currents = this.currents[:len(this.currents)-1]
		return num
	}

	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if len(this.currents) > 0 {
		return this.currents[len(this.currents)-1]
	}

	this.currents = append(this.currents, this.iter.next())
	return this.currents[len(this.currents)-1]
}

// @lc code=end
