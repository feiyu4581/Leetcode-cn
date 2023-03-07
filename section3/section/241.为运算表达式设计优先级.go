package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start

func recursiveDiffWaysToCompute(expression string, index int, operators []byte, value []int) []int {
	numbers := make([]byte, 0)

	collectValue := func() {
		if len(operators) >= 1 && len(value) >= 2 {
			operator := operators[len(operators)-1]
			left, right := value[len(value)-2], value[len(value)-1]

			operators = operators[:len(operators)-1]
			value = value[:len(value)-2]

			switch operator {
			case '+':
				value = append(value, left+right)
			case '-':
				value = append(value, left-right)
			case '*':
				value = append(value, left*right)
			}
		}
	}

	collectNumber := func() {
		if len(numbers) > 0 {
			num, _ := strconv.Atoi(string(numbers))
			numbers = make([]byte, 0)
			value = append(value, num)
		}
	}

	copyInt := func(source []int) []int {
		res := make([]int, 0, len(source))
		for i := 0; i < len(source); i++ {
			res = append(res, source[i])
		}

		return res
	}

	copyByte := func(source []byte) []byte {
		res := make([]byte, 0, len(source))
		for i := 0; i < len(source); i++ {
			res = append(res, source[i])
		}

		return res
	}

	results := make([]int, 0)
	collectOperator := false
	for i := index; i < len(expression); i++ {
		switch expression[i] {
		case '+', '-', '*':
			collectNumber()
			for len(operators) >= 1 && len(value) >= 2 && collectOperator {
				results = append(results, recursiveDiffWaysToCompute(expression, i, copyByte(operators), copyInt(value))...)
				collectValue()
			}

			operators = append(operators, expression[i])
			collectOperator = true
		default:
			numbers = append(numbers, expression[i])
		}
	}

	collectNumber()

	for len(operators) >= 1 {
		collectValue()
	}

	results = append(results, value[0])
	return results
}

func diffWaysToCompute(expression string) []int {
	return recursiveDiffWaysToCompute(expression, 0, make([]byte, 0), make([]int, 0))
}

func Test241() {
	// [0,2]
	fmt.Println(diffWaysToCompute("2-1-1"))

	// [-34,-14,-10,-10,10]
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

// @lc code=end
