package section

import "fmt"

/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 */

// @lc code=start

type Block432 struct {
	prev  *Block432
	next  *Block432
	count int

	items  []*Item432
	cap    int
	length int
}

func (block *Block432) DeleteItem(item *Item432) {
	block.length -= 1
	if item.index != block.length {
		block.items[item.index] = block.items[block.length]
	}
}

func (block *Block432) IsEmpty() bool {
	return block.length == 0
}

func (block *Block432) CheckEmpty() {
	if block.IsEmpty() {
		if block.prev != nil {
			block.prev.next = block.next
		}

		if block.next != nil {
			block.next.prev = block.prev
		}
	}
}

func (block *Block432) AppendItem(item *Item432) {
	if block.length == block.cap {
		block.items = append(block.items, item)
		block.cap += 1
	} else {
		block.items[block.length] = item
	}
	item.index = block.length
	item.block = block
	block.length += 1
}

func (block *Block432) TransferItem(target *Block432, item *Item432) {
	block.DeleteItem(item)
	target.AppendItem(item)
}

func (block *Block432) GetNextBlock() *Block432 {
	newBlock := block.next
	if newBlock == nil || newBlock.count != block.count+1 {
		newBlock = &Block432{
			prev:  block,
			next:  block.next,
			count: block.count + 1,
		}

		if block.next != nil {
			block.next.prev = newBlock
		}
		block.next = newBlock
	}

	return newBlock
}

func (block *Block432) GetPrevBlock() *Block432 {
	newBlock := block.prev
	if newBlock == nil || newBlock.count != block.count-1 {
		newBlock = &Block432{
			prev:  block.prev,
			next:  block,
			count: block.count - 1,
		}

		if block.prev != nil {
			block.prev.next = newBlock
		}
		block.prev = newBlock
	}

	return newBlock
}

func (block *Block432) ChooseOne() string {
	if block != nil && block.length > 0 {
		return block.items[0].key
	}

	return ""
}

type Item432 struct {
	key   string
	count int
	block *Block432
	index int
}

type AllOne struct {
	Head *Block432
	Tail *Block432

	KeyMaps map[string]*Item432
}

func Constructor432() AllOne {
	return AllOne{
		KeyMaps: make(map[string]*Item432),
	}
}

func (this *AllOne) Inc(key string) {
	if item, ok := this.KeyMaps[key]; ok {
		item.count += 1
		block := item.block
		newBlock := block.GetNextBlock()
		if this.Tail == block {
			this.Tail = newBlock
		}
		block.TransferItem(newBlock, item)
		if this.Head == block && block.IsEmpty() {
			this.Head = block.next
		}

		block.CheckEmpty()
	} else {
		item = &Item432{
			key:   key,
			count: 1,
		}

		block := this.Head
		if block == nil || block.count != 1 {
			block = &Block432{
				next:  this.Head,
				count: 1,
			}
			if this.Head != nil {
				this.Head.prev = block
			}
		}
		block.AppendItem(item)
		this.Head = block
		if this.Tail == nil {
			this.Tail = block
		}
		this.KeyMaps[key] = item
	}
}

func (this *AllOne) Dec(key string) {
	if item, ok := this.KeyMaps[key]; ok {
		item.count -= 1
		if item.count == 0 {
			item.block.DeleteItem(item)
			if item.block.IsEmpty() {
				if this.Head == item.block {
					this.Head = item.block.next
				}

				if this.Tail == item.block {
					this.Tail = item.block.prev
				}
				item.block.CheckEmpty()
			}
			delete(this.KeyMaps, key)
			return
		}

		block := item.block
		newBlock := block.GetPrevBlock()
		if this.Head == block {
			this.Head = newBlock
		}

		block.TransferItem(newBlock, item)
		if this.Tail == block && block.IsEmpty() {
			this.Tail = block.prev
		}

		block.CheckEmpty()
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.Tail.ChooseOne()
}

func (this *AllOne) GetMinKey() string {
	return this.Head.ChooseOne()
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

func Test432() {
	allOne := Constructor432()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("b")
	fmt.Println(allOne.GetMaxKey() == "b") // 返回 "b"
	fmt.Println(allOne.GetMinKey() == "a") // 返回 "a"
	allOne.Inc("b")
	allOne.Inc("b")
	allOne.Dec("b")
	allOne.Dec("b")
	fmt.Println(allOne.GetMaxKey() == "b") // 返回 "b"
	fmt.Println(allOne.GetMinKey() == "a") // 返回 "a"

	allOne = Constructor432()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("c")
	allOne.Inc("c")

	allOne.Dec("b")
	allOne.Dec("b")
	fmt.Println(allOne.GetMinKey() == "a") // 返回 "a"
	allOne.Dec("a")
	fmt.Println(allOne.GetMaxKey() == "c") // 返回 "a"
	fmt.Println(allOne.GetMinKey() == "c") // 返回 "a"

	allOne = Constructor432()
	allOne.Inc("hello")
	allOne.Inc("world")
	allOne.Inc("hello")
	allOne.Dec("world")
	allOne.Inc("hello")
	allOne.Inc("leet")
	fmt.Println(allOne.GetMaxKey() == "hello") // 返回 "a"
	allOne.Dec("hello")
	allOne.Dec("hello")
	allOne.Dec("hello")
	fmt.Println(allOne.GetMaxKey() == "leet") // 返回 "a"
}
