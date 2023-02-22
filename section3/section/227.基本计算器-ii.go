package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */
// @lc code=start

type priorityOperator struct {
	operator byte
	priority int
}

var operatorMaps = map[byte]*priorityOperator{
	'+': {
		operator: '+',
		priority: 1,
	},
	'-': {
		operator: '-',
		priority: 1,
	},
	'/': {
		operator: '/',
		priority: 2,
	},
	'*': {
		operator: '*',
		priority: 2,
	},
}

func calculate227(s string) int {
	// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
	numbers := make([]int, 0)
	operators := make([]*priorityOperator, 0)
	currentNumbers := make([]byte, 0)

	isdigit := func(num byte) bool {
		return num >= '0' && num <= '9'
	}

	collectNumber := func() {
		if len(currentNumbers) > 0 {
			number, _ := strconv.Atoi(string(currentNumbers))
			numbers = append(numbers, number)
			currentNumbers = make([]byte, 0)
		}
	}

	handleOperator := func() {
		if len(numbers) >= 2 && len(operators) >= 1 {
			left, right := numbers[len(numbers)-2], numbers[len(numbers)-1]
			nextNumber := 0
			switch operators[len(operators)-1].operator {
			case '+':
				nextNumber = left + right
			case '-':
				nextNumber = left - right
			case '*':
				nextNumber = left * right
			case '/':
				nextNumber = left / right
			}

			numbers = numbers[:len(numbers)-2]
			numbers = append(numbers, nextNumber)
			operators = operators[:len(operators)-1]
		}
	}

	for i := 0; i < len(s); i++ {
		if isdigit(s[i]) {
			currentNumbers = append(currentNumbers, s[i])
		} else {
			collectNumber()
			switch s[i] {
			case '+', '-', '/', '*':
				operator := operatorMaps[s[i]]
				for len(operators) > 0 && operators[len(operators)-1].priority >= operator.priority {
					handleOperator()
				}

				operators = append(operators, operator)
			}
		}
	}

	collectNumber()
	for len(operators) > 0 {
		handleOperator()
	}
	return numbers[0]
}

func Test227() {
	fmt.Println(calculate227("3+2*2") == 7)
	fmt.Println(calculate227(" 3/2 ") == 1)
	fmt.Println(calculate227(" 3+5 / 2 ") == 5)
	fmt.Println(calculate227("1*2-3/4+5*6-7*8+9/10") == -24)
}

// @lc code=end
