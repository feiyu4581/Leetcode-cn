package section

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start

type item381 struct {
	Value int
	Index int
}

type RandomizedCollection struct {
	Items []*item381
	Set   map[int][]*item381
}

func Constructor381() RandomizedCollection {
	return RandomizedCollection{
		Items: make([]*item381, 0),
		Set:   make(map[int][]*item381),
	}
}

func (this *RandomizedCollection) Insert(val int) (res bool) {
	res = false
	if _, ok := this.Set[val]; !ok {
		this.Set[val] = make([]*item381, 0, 1)
		res = true
	}

	item := &item381{
		Value: val,
		Index: len(this.Items),
	}
	this.Set[val] = append(this.Set[val], item)
	this.Items = append(this.Items, item)

	return
}

func (this *RandomizedCollection) Remove(val int) bool {
	if items, ok := this.Set[val]; ok {
		item := items[len(items)-1]

		this.Items[item.Index] = this.Items[len(this.Items)-1]
		this.Items[item.Index].Index = item.Index

		this.Items = this.Items[:len(this.Items)-1]
		if len(items) == 1 {
			delete(this.Set, val)
		} else {
			this.Set[val] = items[:len(items)-1]
		}

		return true
	}

	return false
}

func (this *RandomizedCollection) GetRandom() int {
	return this.Items[rand.Intn(len(this.Items))].Value
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

func Test381() {
	collection := Constructor381()
	fmt.Println(collection.Insert(1))
	fmt.Println(collection.Insert(1))
	fmt.Println(collection.Insert(2))
	fmt.Println(collection.Insert(2))
	fmt.Println(collection.Insert(2))
	fmt.Println(collection.Remove(1))
	fmt.Println(collection.Remove(1))
	fmt.Println(collection.Remove(2))
	fmt.Println(collection.Insert(1))
	fmt.Println(collection.Remove(2))
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
}
