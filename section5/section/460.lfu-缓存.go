package section

import "fmt"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start

type LFUItem struct {
	Key   int
	Value int
	Count int
	Next  *LFUItem
	Prev  *LFUItem
}

type LFULinkedList struct {
	Head *LFUItem
	Tail *LFUItem
}

type LFUCache struct {
	Capacity      int
	KeyMaps       map[int]*LFUItem
	CountHeadMaps map[int]*LFULinkedList
	MinCount      int
}

func Constructor460(capacity int) LFUCache {
	return LFUCache{
		Capacity:      capacity,
		KeyMaps:       make(map[int]*LFUItem, capacity),
		CountHeadMaps: make(map[int]*LFULinkedList, capacity),
	}
}

func (this *LFUCache) appendItem(item *LFUItem) {
	item.Prev, item.Next = nil, nil

	if this.CountHeadMaps[item.Count] == nil {
		this.CountHeadMaps[item.Count] = &LFULinkedList{
			Head: item,
			Tail: item,
		}
	} else {
		linkedList := this.CountHeadMaps[item.Count]
		if linkedList.Tail == nil {
			linkedList.Head, linkedList.Tail = item, item
		} else {
			linkedList.Tail.Next = item
			item.Prev = linkedList.Tail

			linkedList.Tail = item
		}
	}
}

func (this *LFUCache) upgradeCount(item *LFUItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	}
	if item.Next != nil {
		item.Next.Prev = item.Prev
	}

	if linkedList := this.CountHeadMaps[item.Count]; linkedList != nil {
		if linkedList.Head == linkedList.Tail {
			linkedList.Head, linkedList.Tail = nil, nil
		} else if linkedList.Head == item {
			linkedList.Head = item.Next
		} else if linkedList.Tail == item {
			linkedList.Tail = item.Prev
		}
	}
	item.Next, item.Prev = nil, nil

	if this.CountHeadMaps[item.Count].Head == nil && item.Count == this.MinCount {
		this.MinCount += 1
	}

	item.Count += 1
	this.appendItem(item)
}

func (this *LFUCache) Get(key int) int {
	if item, ok := this.KeyMaps[key]; ok {
		value := item.Value
		this.upgradeCount(item)
		return value
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if item, ok := this.KeyMaps[key]; ok {
		item.Value = value
		this.upgradeCount(item)
		return
	}

	if len(this.KeyMaps) == this.Capacity {
		linkedList := this.CountHeadMaps[this.MinCount]
		lastItem := linkedList.Head
		linkedList.Head = lastItem.Next

		delete(this.KeyMaps, lastItem.Key)

		if linkedList.Head == nil {
			linkedList.Tail = nil
			for len(this.KeyMaps) > 0 && (this.CountHeadMaps[this.MinCount] == nil || this.CountHeadMaps[this.MinCount].Head == nil) {
				this.MinCount += 1
			}
		}
	}

	item := &LFUItem{
		Key:   key,
		Value: value,
		Count: 1,
	}

	this.KeyMaps[key] = item
	this.appendItem(item)
	this.MinCount = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func Test460() {
	lfu := Constructor460(2)
	lfu.Put(1, 1)                // cache=[1,_], cnt(1)=1
	lfu.Put(2, 2)                // cache=[2,1], cnt(2)=1, cnt(1)=1
	fmt.Println(lfu.Get(1) == 1) // 返回 1
	// cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.Put(3, 3) // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
	// cache=[3,1], cnt(3)=1, cnt(1)=2
	fmt.Println(lfu.Get(2) == -1) // 返回 -1（未找到）
	fmt.Println(lfu.Get(3) == 3)  // 返回 3
	// cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.Put(4, 4) // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	fmt.Println(lfu.Get(1) == -1) // 返回 -1（未找到）
	fmt.Println(lfu.Get(3) == 3)  // 返回 3
	// cache=[3,4], cnt(4)=1, cnt(3)=3
	fmt.Println(lfu.Get(4) == 4) // 返回 4
	// cache=[3,4], cnt(4)=2, cnt(3)=3

	fmt.Println("----new---")
	lfu = Constructor460(1)
	lfu.Put(2, 1)
	fmt.Println(lfu.Get(2) == 1) // 返回 1
	lfu.Put(3, 2)
	fmt.Println(lfu.Get(2) == -1) // 返回 1
	fmt.Println(lfu.Get(3) == 2)  // 返回 1
}
