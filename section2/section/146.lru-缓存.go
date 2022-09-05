package section

import "fmt"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start

type cacheItem struct {
	Key   int
	Value int
	Next  *cacheItem
	Prev  *cacheItem
}

type LRUCache struct {
	Head   *cacheItem
	Tail   *cacheItem
	Caches map[int]*cacheItem
	Size   int
	Length int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Head:   nil,
		Tail:   nil,
		Caches: make(map[int]*cacheItem, capacity),
		Size:   capacity,
		Length: 0,
	}
}

func (this *LRUCache) putTail(item *cacheItem) {
	item.Next, item.Prev = nil, nil
	if this.Head == nil {
		this.Head = item
		this.Tail = item
	} else {
		this.Tail.Next = item
		item.Prev = this.Tail
		this.Tail = item
	}
}

func (this *LRUCache) moveToTail(item *cacheItem) {
	if this.Tail != item {
		if this.Head == item {
			this.Head = item.Next
		}
		if item.Prev != nil {
			item.Prev.Next = item.Next
		}
		if item.Next != nil {
			item.Next.Prev = item.Prev
		}
		this.putTail(item)
	}
}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.Caches[key]; ok {
		this.moveToTail(item)
		return item.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if item, ok := this.Caches[key]; ok {
		item.Value = value
		this.moveToTail(item)
	} else {
		item = &cacheItem{
			Key:   key,
			Value: value,
		}
		this.moveToTail(item)
		if this.Length == this.Size {
			delete(this.Caches, this.Head.Key)
			this.Head = this.Head.Next
			this.Head.Prev = nil
		} else {
			this.Length += 1
		}
		this.Caches[key] = item
	}
}

func Test146() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)                 // 缓存是 {1=1}
	lRUCache.Put(2, 2)                 // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1) == 1)  // 返回 1
	lRUCache.Put(3, 3)                 // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2) == -1) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)                 // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1) == -1) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3) == 3)  // 返回 3
	fmt.Println(lRUCache.Get(4) == 4)  // 返回 4
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
