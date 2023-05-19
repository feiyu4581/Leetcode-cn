package section

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start

type Set380 struct {
	Index int
	Value int
}

type RandomizedSet struct {
	set   map[int]*Set380
	slice []*Set380
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]*Set380),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}

	this.set[val] = &Set380{
		Index: len(this.slice),
		Value: val,
	}

	this.slice = append(this.slice, this.set[val])

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if set, ok := this.set[val]; ok {
		this.slice[set.Index] = this.slice[len(this.slice)-1]
		this.slice[set.Index].Index = set.Index
		this.slice = this.slice[:len(this.slice)-1]

		delete(this.set, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.slice[rand.Intn(len(this.slice))].Value
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

func Test380() {
	randomizedSet := Constructor380()
	fmt.Println(randomizedSet.Insert(0) == true) // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	fmt.Println(randomizedSet.Insert(1) == true) // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	fmt.Println(randomizedSet.Remove(0) == true) // 返回 false ，表示集合中不存在 2 。
	fmt.Println(randomizedSet.Insert(2) == true) // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	fmt.Println(randomizedSet.Remove(1) == true) // 返回 false ，表示集合中不存在 2 。
	fmt.Println(randomizedSet.GetRandom())       // getRandom 应随机返回 1 或 2 。
}
