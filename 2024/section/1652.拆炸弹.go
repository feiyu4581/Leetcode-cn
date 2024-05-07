package section

import "fmt"

/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		return res
	}

	count := k
	if count < 0 {
		count = -count
	}

	current := 0
	for i := 0; i < count; i++ {
		current += code[i]
	}

	next := func(index int) int {
		if nextIndex := index + 1; nextIndex >= len(code) {
			return nextIndex % len(code)
		} else {
			return nextIndex
		}
	}

	prev := func(index int) int {
		if prevIndex := index - 1; prevIndex == -1 {
			return len(code) - 1
		} else {
			return prevIndex
		}
	}

	for i := 0; i < len(code); i++ {
		if k > 0 {
			res[prev(i)] = current
		} else {
			res[next(i+count-1)] = current
		}

		current = current - code[i] + code[next(i+count-1)]
	}

	return res
}

// @lc code=end

func Test1652() {
	// [12,10,16,13]
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	// [0,0,0,0]
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))
	// [12,5,6,13]
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}
