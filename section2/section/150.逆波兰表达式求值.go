package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stacks := make([]int, 0, len(tokens))
	popItems := func() (int, int) {
		defer func() {
			stacks = stacks[0 : len(stacks)-2]
		}()
		return stacks[len(stacks)-2], stacks[len(stacks)-1]
	}
	for _, token := range tokens {
		switch token {
		case "+":
			left, right := popItems()
			stacks = append(stacks, left+right)
		case "-":
			left, right := popItems()
			stacks = append(stacks, left-right)
		case "*":
			left, right := popItems()
			stacks = append(stacks, left*right)
		case "/":
			left, right := popItems()
			stacks = append(stacks, left/right)
		default:
			num, _ := strconv.Atoi(token)
			stacks = append(stacks, num)
		}
	}

	return stacks[0]
}

func Test150() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}) == 9)
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}) == 6)
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}) == 22)
}

// @lc code=end
